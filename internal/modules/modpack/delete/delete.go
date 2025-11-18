package delete

import (
	"context"

	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/domain"
)

type ModpackDelete interface {
	Delete(ctx context.Context, srv domain.ServerID, id domain.ModpackID) error
}
