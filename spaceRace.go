package main

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"spacerace/entity"
	"spacerace/systems"
)

type mainScene struct {
}

// Type uniquely defines your game type
func (m *mainScene) Type() string { return "myGame" }

// Preload is called before loading any assets from the disk,
// to allow you to register / queue them
func (m *mainScene) Preload() {
	err := engo.Files.Load(
		"textures/playerShip1_blue.png",
		"textures/playerShip1_green.png",
		"textures/meteorBrown_med1.png",
	)
	if err != nil {
		return
	}
}

// Setup is called before the main loop starts. It allows you
// to add entities and systems to your Scene.
func (m *mainScene) Setup(u engo.Updater) {
	engo.Input.RegisterButton("A", engo.KeyW)
	engo.Input.RegisterButton("Up", engo.KeyArrowUp)

	world, _ := u.(*ecs.World)
	world.AddSystem(&common.RenderSystem{})

	world.AddSystem(&systems.PlayerControlSystem{})
	r := systems.RockSpawnSystemSystem{}
	world.AddSystem(&r)

	spaceShipLeft := entity.SpaceShip{BasicEntity: ecs.NewBasic(), PlayerId: 1}
	spaceShipRight := entity.SpaceShip{BasicEntity: ecs.NewBasic(), PlayerId: 2}
	spaceShipLeft.SpaceComponent = common.SpaceComponent{
		Width:  99,
		Height: 75,
	}
	spaceShipLeft.Position = engo.Point{X: (engo.GameWidth() / 2) - (spaceShipLeft.Width + 200/2), Y: engo.GameHeight() - spaceShipLeft.Height - 20}

	spaceShipRight.SpaceComponent = common.SpaceComponent{
		Width:  99,
		Height: 75,
	}
	spaceShipRight.Position = engo.Point{X: (engo.GameWidth() / 2) + (spaceShipRight.Width + 50/2), Y: engo.GameHeight() - spaceShipRight.Height - 20}

	spaceShipLeftTexture, err := common.LoadedSprite("textures/playerShip1_blue.png")
	if err != nil {
		return
	}

	spaceShipRightTexture, err := common.LoadedSprite("textures/playerShip1_green.png")
	if err != nil {
		return
	}

	spaceShipRight.RenderComponent = common.RenderComponent{
		Drawable: spaceShipRightTexture,
		Scale:    engo.Point{X: 1.0, Y: 1.0},
	}
	spaceShipLeft.RenderComponent = common.RenderComponent{
		Drawable: spaceShipLeftTexture,
		Scale:    engo.Point{X: 1.0, Y: 1.0},
	}

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&spaceShipLeft.BasicEntity, &spaceShipLeft.RenderComponent, &spaceShipLeft.SpaceComponent)
			sys.Add(&spaceShipRight.BasicEntity, &spaceShipRight.RenderComponent, &spaceShipRight.SpaceComponent)
		case *systems.PlayerControlSystem:
			sys.Add(&spaceShipLeft.BasicEntity, &spaceShipLeft.SpaceComponent, &spaceShipLeft.PlayerId)
			sys.Add(&spaceShipRight.BasicEntity, &spaceShipRight.SpaceComponent, &spaceShipRight.PlayerId)
		case *systems.RockSpawnSystemSystem:
		}
	}

	rockBrown1Texture, err := common.LoadedSprite("textures/meteorBrown_med1.png")
	if err != nil {
		panic(err)
	}

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&spaceShipLeft.BasicEntity, &spaceShipLeft.RenderComponent, &spaceShipLeft.SpaceComponent)
			sys.Add(&spaceShipRight.BasicEntity, &spaceShipRight.RenderComponent, &spaceShipRight.SpaceComponent)
		case *systems.PlayerControlSystem:
			sys.Add(&spaceShipLeft.BasicEntity, &spaceShipLeft.SpaceComponent, &spaceShipLeft.PlayerId)
			sys.Add(&spaceShipRight.BasicEntity, &spaceShipRight.SpaceComponent, &spaceShipRight.PlayerId)
		case *systems.RockSpawnSystemSystem:
			sys.Add(rockBrown1Texture)
		}
	}

}

func main() {
	opts := engo.RunOptions{
		Title:        "Space Race",
		Width:        800,
		Height:       1000,
		NotResizable: true,
	}
	engo.Run(opts, &mainScene{})
}
