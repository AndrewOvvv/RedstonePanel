package install

import (
	"context"

	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/domain"
)

type PluginInstall interface {
	Install(ctx context.Context, srv domain.ServerID, plugin domain.PluginRef) error
}
