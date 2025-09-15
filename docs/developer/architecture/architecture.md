# RedstonePanel Architecture

**RedstonePanel** is built upon a **modular, message-driven architecture**. This design was chosen to achieve the following key goals:

-   **Loose Coupling:** Modules have no direct dependencies on one another. This allows us to modify, replace, or remove a module without impacting the rest of the application.
-   **High Testability:** Every module can be tested in isolation by providing mock implementations for its external dependencies.
-   **Scalability & Extensibility:** Adding new functionality is as simple as creating a new module that communicates with the system through standard, well-defined contracts.

## Core Entities

The architecture is composed of two fundamental entities:

### 1. [Module](developer/architecture/module.md)

A **Module** is an isolated component that encapsulates specific business logic. Each module is responsible for a well-defined domain. For example:
-   `servers` — Managing the game server lifecycle.
-   `plugins` — Installing and configuring plugins.
-   `users` — Handling user authentication and authorization.

**The Golden Rule:** A module must have no direct dependencies on other modules and must not know about their internal implementation.

### 2. [Kernel](developer/architecture/kernel.md)

The **Kernel** is the central Message Bus that acts as the mediator for all communication between modules in RedstonePanel. The Kernel itself contains no business logic; its sole responsibility is to route two types of messages: **Commands** and **Events**.

## Communication Model

All communication between modules happens **exclusively through the Kernel**. Direct inter-module calls are strictly forbidden.

This communication is based on two distinct patterns:

#### Commands
-   **Intent:** "Do this" (e.g., `CreateUser`, `StopServer`).
-   **Handler:** Exactly **one** handler exists for a command across the entire application.
-   **Type:** **Synchronous** request-response. The caller dispatches the command and blocks until it receives a result or an error.

#### Events
-   **Intent:** "This happened" (e.g., `UserCreated`, `ServerStopped`).
-   **Listeners:** **Zero, one, or many** listeners can subscribe to an event.
-   **Type:** **Asynchronous** "fire-and-forget". The publisher dispatches an event and does not wait for it to be processed.

## How It Works Together

1.  At application startup, instances of all **Modules** and the **Kernel** are created.
2.  Each Module **registers** its Command handlers and Event listeners with the Kernel.
3.  When Module A needs Module B to perform an action, it **sends** a typed **Command** to the Kernel.
4.  The Kernel finds the single registered handler for that Command (in Module B) and **invokes** it, awaiting a response.
5.  The handler in Module B executes its business logic. Upon completion, it may **publish** one or more **Events** to the Kernel.
6.  The Kernel finds all subscribers for that Event and **asynchronously notifies** them.

This model ensures **RedstonePanel** remains flexible and predictable. Modules do not communicate with each other, but rather with the application as a whole through well-defined messages.

For detailed information on the Kernel interface, usage examples, and Module structure, please refer to the respective documentation pages.
