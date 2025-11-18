package list

import (
	"context"

	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/domain"
)

type PluginList interface {
	List(ctx context.Context, srv domain.ServerID) ([]domain.PluginRef, error)
}
