package delete

import (
	"context"

	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/domain"
)

type PluginDelete interface {
	Delete(ctx context.Context, srv domain.ServerID, pluginID domain.PluginID) error
}
