package install

import (
	"context"

	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/domain"
)

type InstanceInstall interface {
	Install(ctx context.Context, m *domain.Manifest) error
}
