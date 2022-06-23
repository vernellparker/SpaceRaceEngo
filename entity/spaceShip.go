package entity

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
	"spacerace/components"
)

type SpaceShip struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	components.Moveable
	PlayerId int
}
