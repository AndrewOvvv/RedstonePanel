package plugin

import (
	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/modules/plugin/delete"
	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/modules/plugin/install"
	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/modules/plugin/list"
)

type Plugin interface {
	install.PluginInstall
	delete.PluginDelete
	list.PluginList
}
