# The Kernel: A Decoupled Messaging Architecture

This document describes the core architectural component of our application: the **Kernel**. It serves as a central message bus, enabling communication between different modules in a decoupled, scalable, and maintainable way.

The fundamental principle of this architecture is: **Modules must not know about each other.** All interaction is mediated by the Kernel through a clear system of Commands and Events.

## Core Concepts

There are two primary ways modules communicate through the Kernel: **Commands** and **Events**. Understanding the difference is crucial.

### Commands
A Command is an imperative message that represents a request to perform a specific action.

-   **Intent:** "Do this thing." (e.g., `CreateUser`, `StartServer`)
-   **Handler:** A command has **exactly one** handler in the entire system.
-   **Communication:** **Synchronous**. The caller sends the command and a-waits for a response or an error. This is akin to a remote function call.
-   **Usage:** Use a Command when the caller needs a direct result from the action or needs to know if it succeeded immediately.

### Events
An Event is a declarative message that signifies something has already happened.

-   **Intent:** "This thing happened." (e.g., `UserCreated`, `ServerStarted`)
-   **Listeners:** An event can have **zero, one, or many** listeners.
-   **Communication:** **Asynchronous**. The publisher sends the event and does not wait for the listeners to complete. This is a "fire-and-forget" mechanism.
-   **Usage:** Use an Event to notify other parts of the system about a change in state, allowing them to react independently.

| Feature         | Command                                       | Event                                          |
| --------------- | --------------------------------------------- | ---------------------------------------------- |
| **Philosophy**  | Tell a specific module to do something        | Announce to the system that something happened |
| **Handlers**    | Exactly one                                   | Zero or more                                   |
| **Execution**   | Synchronous (Request-Response)                | Asynchronous (Fire-and-Forget)                 |
| **Coupling**    | Caller depends on the Command, not the handler | Publisher and listeners are completely decoupled |
| **Return Value**| Must return a response and/or an `error`      | No direct return value to the publisher        |

---

## The `Kernel` Interface

This is the public API of the Kernel, designed for complete type safety using Go 1.18+ Generics.

```go
package kernel

// The Kernel is the central mediator for all inter-module communication.
type Kernel interface {
    // Execute dispatches a command to its single handler and returns a response.
    // It is a synchronous, blocking call.
    // C is the command type, R is the response type.
    Execute[C any, R any](ctx context.Context, command C) (R, error)

    // Publish dispatches an event to all registered listeners.
    // It is an asynchronous, non-blocking call from the publisher's perspective.
    // E is the event type.
    // The returned error only indicates an immediate problem dispatching (e.g., kernel is closed),
    // not errors from the asynchronous listeners.
    Publish[E any](ctx context.Context, event E) error

    // RegisterCommandHandler registers a function as the single handler for a command type.
    // The command and response types (C and R) are inferred from the handler's signature.
    // Panics if a handler for the same command type is already registered.
    RegisterCommandHandler[C any, R any](handler func(ctx context.Context, cmd C) (R, error))

    // RegisterEventHandler registers a function as a listener for an event type.
    // The event type (E) is inferred from the handler's signature.
    // Multiple listeners can be registered for the same event type.
    RegisterEventHandler[E any](handler func(ctx context.Context, event E) error)
}
```

---

## Full Usage Example (just example, for realization check source code)

Let's model a `servers` module that manages game servers and a `notifications` module that sends alerts.

### 1. Working with a Command: `StopServer`

#### Step 1: Define the Command and its Response

These are simple data structures, often placed in a shared package or within the module that "owns" the business logic.

```go
// in file: modules/servers/contracts.go

package servers

// StopServerCommand is the request to stop a server.
type StopServerCommand struct {
    ServerID string
    Reason   string
}

// StopServerResponse is the result of a successful stop operation.
type StopServerResponse struct {
    ServerID  string
    FinalStatus string
}
```

#### Step 2: Implement the Command Handler

In your module, create a method that matches the required signature. It's fully type-safe—no `any` or type assertions needed!

```go
// in file: modules/servers/module.go

package servers

import "context"

type Module struct {
    // ... module's dependencies, like a logger and etc. 
}

// HandleStopServer is the concrete implementation for stopping a server.
// Note the type-safe signature: it directly accepts StopServerCommand.
func (m *Module) HandleStopServer(ctx context.Context, cmd StopServerCommand) (StopServerResponse, error) {
    log.Printf(ctx, "Received request to stop server %s: %s", cmd.ServerID, cmd.Reason)
    
    // ... actual logic to stop the server ...
    
    // After stopping, we create and return a typed response.
    response := StopServerResponse{
        ServerID:  cmd.ServerID,
        FinalStatus: "stopped",
    }
    
    return response, nil
}
```

#### Step 3: Register the Handler

During application startup (e.g., in `main.go`), you wire everything together.

```go
// in file: cmd/app/main.go

import (
    "myapp/kernel"
    "myapp/modules/servers"
)

func main() {
    theKernel := kernel.New() // Create a new kernel instance
    serverModule := servers.NewModule(/* dependencies */)
    
    // The Go compiler infers the generic types C and R automatically!
    theKernel.RegisterCommandHandler(serverModule.HandleStopServer)
    
    // ... start the application
}
```

#### Step 4: Execute the Command

Now, any other part of the application (e.g., an HTTP handler) can execute this command without knowing anything about the `servers` module.

```go
// in file: api/http_handlers.go

func stopServerHandler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    
    cmd := servers.StopServerCommand{
        ServerID: "mc-creative-1",
        Reason:   "Scheduled restart via API",
    }
    
    // Execute the command. Both `response` and `err` are fully typed.
    response, err := theKernel.Execute[servers.StopServerCommand, servers.StopServerResponse](ctx, cmd)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    
    fmt.Fprintf(w, "Server %s is now %s", response.ServerID, response.FinalStatus)
}
```

### 2. Working with an Event: `ServerStopped`

#### Step 1: Define the Event

The `servers` module can publish an event after it has successfully stopped a server.

```go
// in file: modules/servers/contracts.go

package servers

// ServerStoppedEvent is published when a server has fully stopped.
type ServerStoppedEvent struct {
    ServerID  string
    StoppedAt time.Time
}
```

#### Step 2: Implement Event Listeners

Multiple modules might be interested in this event.

```go
// in file: modules/notifications/module.go

// HandleServerStopped sends a notification to an admin.
func (m *NotificationsModule) HandleServerStopped(ctx context.Context, event servers.ServerStoppedEvent) error {
    msg := fmt.Sprintf("Alert: Server %s was stopped at %s", event.ServerID, event.StoppedAt)
    return m.slackClient.SendMessage(ctx, "#ops-alerts", msg)
}

// in file: modules/auditing/module.go

// LogServerStopped writes a record to an audit log.
func (m *AuditingModule) LogServerStopped(ctx context.Context, event servers.ServerStoppedEvent) error {
    return m.auditLog.Record(ctx, "server.stopped", event)
}
```

#### Step 3: Register Listeners

In your startup code, register all interested listeners.

```go
// in file: cmd/app/main.go

func main() {
    // ... (setup from before)
    notificationsModule := notifications.NewModule(/* ... */)
    auditingModule := auditing.NewModule(/* ... */)
    
    theKernel.RegisterEventHandler(notificationsModule.HandleServerStopped)
    theKernel.RegisterEventHandler(auditingModule.LogServerStopped)
    
    // ...
}
```

#### Step 4: Publish the Event

The command handler is the perfect place to publish the corresponding event.

```go
// in file: modules/servers/module.go

func (m *Module) HandleStopServer(ctx context.Context, cmd StopServerCommand) (StopServerResponse, error) {
    // ... logic to stop the server ...
    
    // On success, publish an event for other modules to consume.
    event := ServerStoppedEvent{
        ServerID:  cmd.ServerID,
        StoppedAt: time.Now(),
    }
    // This is a fire-and-forget call.
    _ = m.kernel.Publish(ctx, event) // Assuming module has a reference to the kernel
    
    response := StopServerResponse{ /* ... */ }
    return response, nil
}
```
