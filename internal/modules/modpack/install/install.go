package install

import (
	"context"

	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/domain"
)

type ModpackInstall interface {
	Install(ctx context.Context, srv domain.ServerID, mp domain.ModpackRef) error
}
