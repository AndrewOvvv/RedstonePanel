package install

import (
	"context"

	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/domain"
)

type ModInstall interface {
	Install(ctx context.Context, srv domain.ServerID, mod domain.ModRef) error
}
