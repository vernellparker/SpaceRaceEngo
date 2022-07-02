package entity

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

type PlayerScore struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}
