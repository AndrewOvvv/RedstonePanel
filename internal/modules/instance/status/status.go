package status

import (
	"context"

	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/domain"
)

type Status struct {
	Running   bool
	Pid       int64
	ListenTCP []domain.ListenTCP
}

type InstanceStatus interface {
	Status(ctx context.Context, m domain.ServerID) (Status, error)
}
