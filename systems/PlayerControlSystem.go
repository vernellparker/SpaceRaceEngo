package systems

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type playerControlSystemEntity struct {
	*ecs.BasicEntity
	*common.SpaceComponent
	*common.RenderComponent
	playerId *int
}
type PlayerControlSystem struct {
	world    *ecs.World
	entities []playerControlSystemEntity
}

func (c *PlayerControlSystem) New(w *ecs.World) {
	c.world = w
}

func (c *PlayerControlSystem) Add(b *ecs.BasicEntity, s *common.SpaceComponent, p *int) {
	c.entities = append(c.entities, playerControlSystemEntity{
		BasicEntity:    b,
		SpaceComponent: s,
		playerId:       p,
	})
}

func (c *PlayerControlSystem) Update(dt float32) {
	var moveSpeed float32 = 5.0

	for _, e := range c.entities {
		if btn1 := engo.Input.Button("A"); btn1.Down() {
			if *e.playerId == 1 {
				e.Position = engo.Point{X: e.Position.X, Y: e.Position.Y - moveSpeed}
			}
		}

		if btn2 := engo.Input.Button("Up"); btn2.Down() {
			if *e.playerId == 2 {
				e.Position = engo.Point{X: e.Position.X, Y: e.Position.Y - moveSpeed}
			}
		}
	}
}

func (c *PlayerControlSystem) Remove(e ecs.BasicEntity) {
	delete := -1
	for index, en := range c.entities {
		if en.BasicEntity.ID() == e.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		c.entities = append(c.entities[:delete], c.entities[delete+1:]...)
	}
}
