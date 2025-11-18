package delete

import (
	"context"

	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/domain"
)

type ModDelete interface {
	Delete(ctx context.Context, srv domain.ServerID, modID domain.ModID) error
}
