package systems

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"math/rand"
	"spacerace/entity"
	"sync"
	"time"
)

var m sync.Mutex

type rockSpawnSystemEntity struct {
	*ecs.BasicEntity
	*common.SpaceComponent
	*common.RenderComponent
	Direction int
}

type RockSpawnSystemSystem struct {
	world       *ecs.World
	rockTexture *common.Texture
	entities    []*rockSpawnSystemEntity
	speed       float32
	xDirection  float32
}

func (r *RockSpawnSystemSystem) New(w *ecs.World) {
	r.world = w
	r.speed = 50
}

func (r *RockSpawnSystemSystem) Add(texture *common.Texture) {
	r.rockTexture = texture
	go func() {
		spawnLeft := true
		counter := 0
		for counter < 10 {
			r.CreateRockEntities(&m, spawnLeft)
			time.Sleep(500 * time.Millisecond)
			counter++
			spawnLeft = !spawnLeft
		}
	}()
}

func (r *RockSpawnSystemSystem) Update(dt float32) {
	for _, rock := range r.entities {
		if rock.Direction == 1 {
			rock.Position.X = rock.Position.X + 1*r.speed*dt
			if rock.Position.X >= engo.GameWidth()+rock.Width {
				rock.Direction = 2
			}
		} else {
			rock.Position.X = rock.Position.X - 1*r.speed*dt
			if rock.Position.X <= 0-(rock.Width+20) {
				rock.Direction = 1
			}
		}

	}
}

func (r *RockSpawnSystemSystem) Remove(e ecs.BasicEntity) {}

func (r *RockSpawnSystemSystem) CreateRockEntities(m *sync.Mutex, left bool) {
	s := rand.NewSource(time.Now().UnixNano())
	rd := rand.New(s)

	rock := entity.Rock{BasicEntity: ecs.NewBasic()}
	rock.RenderComponent = common.RenderComponent{
		Drawable: r.rockTexture,
		Scale:    engo.Point{X: 1.0, Y: 1.0},
	}

	rock.SpaceComponent = common.SpaceComponent{
		Width:  43,
		Height: 43,
	}
	var xPosition float32
	var xDirection int
	if left {
		xPosition = 0 + rock.Width
		xDirection = 1
	} else {
		xPosition = engo.GameWidth() - rock.Width
		xDirection = 2
	}

	rock.Position = engo.Point{X: xPosition,
		Y: float32(rd.Intn(int(engo.GameHeight()-100))) - rock.Height}

	rock.CollisionComponent = common.CollisionComponent{Main: 1}

	for _, system := range r.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			m.Lock()
			sys.Add(&rock.BasicEntity, &rock.RenderComponent, &rock.SpaceComponent)
			m.Unlock()
		case *common.CollisionSystem:
			m.Lock()
			sys.Add(&rock.BasicEntity, &rock.CollisionComponent, &rock.SpaceComponent)
			m.Unlock()

		}

		r.entities = append(r.entities, &rockSpawnSystemEntity{
			&rock.BasicEntity,
			&rock.SpaceComponent,
			&rock.RenderComponent,
			xDirection,
		})

	}
}
