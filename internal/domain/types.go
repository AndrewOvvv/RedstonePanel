package domain

import "time"

type ServerLoader string

const (
	Vanilla ServerLoader = "vanilla"
	Forge   ServerLoader = "forge"
	Fabric  ServerLoader = "fabric"
	Quilt   ServerLoader = "quilt"
	Paper   ServerLoader = "paper"
	Spigot  ServerLoader = "spigot"
	Bukkit  ServerLoader = "bukkit"
	Purpur  ServerLoader = "purpur"
	Mohist  ServerLoader = "mohist"
)

// ServerID uniquely identifies a managed Minecraft server instance.
type ServerID string

// ModID uniquely identifies a Minecraft mod.
// It contains in META-INF/mods.toml.
type ModID string

// PluginID uniquely identifies a server plugin.
// It contains in plugin.yml.
type PluginID string

// ModpackID uniquely identifies a modpack definition/version.
// Concatenation of modpack name and modpack minecraft version.
type ModpackID string

// ManifestID uniquely identifies a manifest stored for a server.
type ManifestID string

// MinecraftVersion represents release numbers.
// For example, 1.21.8.
type MinecraftVersion string

// Path is a filesystem path on the host.
type Path string

// URL stands for unified resource location.
type URL string

type ListenTCP struct {
	IP   string
	Port int
}

// Manifest describes desired state for a server instance.
type Manifest struct {
	ID        ManifestID
	Server    ServerID
	CreatedAt time.Time
	UpdatedAt time.Time

	Loader  ServerLoader
	Version MinecraftVersion

	JavaArgs []string
	Env      map[string]string
	DataDir  Path

	// Optional fields
	Mods    []ModRef
	Plugins []PluginRef
	Modpack *ModpackRef
}

// ModRef pins a mod to install.
type ModRef struct {
	ID      ModID
	Name    string
	Version MinecraftVersion
	Source  URL
}

// PluginRef pins a plugin to install.
type PluginRef struct {
	ID      PluginID
	Name    string
	Version MinecraftVersion
	Source  URL
}

// ModpackRef pins a modpack to install.
type ModpackRef struct {
	ID      ModpackID
	Name    string
	Version MinecraftVersion
	Source  URL
}
