package get

import (
	"context"

	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/domain"
)

type InstanceGet interface {
	Get(ctx context.Context, m domain.ServerID) (*domain.Manifest, error)
}
