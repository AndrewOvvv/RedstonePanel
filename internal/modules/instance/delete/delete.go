package delete

import (
	"context"

	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/domain"
)

type InstanceDelete interface {
	Delete(ctx context.Context, m domain.ServerID) error
}
