package instance

import (
	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/modules/instance/delete"
	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/modules/instance/get"
	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/modules/instance/list"
	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/modules/instance/start"
	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/modules/instance/status"
	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/modules/instance/stop"
)

type Instance interface {
	get.InstanceGet
	list.InstanceList
	start.InstanceStart
	stop.InstanceStop
	status.InstanceStatus
	delete.InstanceDelete
}
