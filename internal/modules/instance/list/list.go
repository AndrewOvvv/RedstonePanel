package list

import (
	"context"

	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/domain"
)

type InstanceList interface {
	List(ctx context.Context) ([]domain.Manifest, error)
}
