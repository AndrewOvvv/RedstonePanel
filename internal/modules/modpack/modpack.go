package modpack

import (
	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/modules/modpack/delete"
	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/modules/modpack/install"
)

type Modpack interface {
	install.ModpackInstall
	delete.ModpackDelete
}
