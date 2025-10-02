package modules

import (
	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/modules/instance"
	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/modules/mod"
	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/modules/modpack"
	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/modules/plugin"
)

type Modules struct {
	Instance instance.Instance
	Mod      mod.Mod
	Plugin   plugin.Plugin
	Modpack  modpack.Modpack
}
