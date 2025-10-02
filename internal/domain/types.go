package domain

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
