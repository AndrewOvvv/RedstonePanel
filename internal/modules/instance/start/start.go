package start

import (
	"context"

	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/domain"
)

type InstanceStart interface {
	Start(ctx context.Context, m domain.ServerID) error
}
