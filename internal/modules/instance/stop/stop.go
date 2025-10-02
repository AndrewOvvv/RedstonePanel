package stop

import (
	"context"

	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/domain"
)

type InstanceStop interface {
	Stop(ctx context.Context, m domain.ServerID) error
}
