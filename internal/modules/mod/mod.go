package mod

import (
	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/modules/mod/delete"
	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/modules/mod/install"
	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/modules/mod/list"
)

type Mod interface {
	list.ModList
	install.ModInstall
	delete.ModDelete
}
