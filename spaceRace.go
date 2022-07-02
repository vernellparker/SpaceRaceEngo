package main

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"image/color"
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
		"fonts/PressStart2P-Regular.ttf",
	)
	if err != nil {
		return
	}
}

// Setup is called before the main loop starts. It allows you
// to add entities and systems to your Scene.
func (m *mainScene) Setup(u engo.Updater) {
	//#region Input
	engo.Input.RegisterButton("A", engo.KeyW)
	engo.Input.RegisterButton("Up", engo.KeyArrowUp)
	//#endregion Input

	//#region Add Systems
	world, _ := u.(*ecs.World)
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&common.CollisionSystem{Solids: 1})

	world.AddSystem(&systems.PlayerControlSystem{})
	r := systems.RockSpawnSystemSystem{}
	world.AddSystem(&r)
	//#endregion Add Systems

	//#region SpaceShip Left
	spaceShipLeft := entity.SpaceShip{BasicEntity: ecs.NewBasic(), PlayerId: 1}
	spaceShipRight := entity.SpaceShip{BasicEntity: ecs.NewBasic(), PlayerId: 2}
	spaceShipLeft.SpaceComponent = common.SpaceComponent{
		Width:  99,
		Height: 75,
	}
	spaceShipLeft.Position = engo.Point{X: (engo.GameWidth() / 2) - (spaceShipLeft.Width + 200/2), Y: engo.GameHeight() - spaceShipLeft.Height - 20}
	spaceShipLeft.CollisionComponent = common.CollisionComponent{
		Group: 1,
	}
	spaceShipLeftTexture, err := common.LoadedSprite("textures/playerShip1_blue.png")
	if err != nil {
		return
	}

	spaceShipLeft.RenderComponent = common.RenderComponent{
		Drawable: spaceShipLeftTexture,
		Scale:    engo.Point{X: 1.0, Y: 1.0},
	}
	//#endregion SpaceShip Left

	//#region SpaceShip Right
	spaceShipRight.SpaceComponent = common.SpaceComponent{
		Width:  99,
		Height: 75,
	}
	spaceShipRight.Position = engo.Point{X: (engo.GameWidth() / 2) + (spaceShipRight.Width + 50/2), Y: engo.GameHeight() - spaceShipRight.Height - 20}
	spaceShipRight.CollisionComponent = common.CollisionComponent{
		Group: 1,
	}

	spaceShipRightTexture, err := common.LoadedSprite("textures/playerShip1_green.png")
	if err != nil {
		return
	}

	spaceShipRight.RenderComponent = common.RenderComponent{
		Drawable: spaceShipRightTexture,
		Scale:    engo.Point{X: 1.0, Y: 1.0},
	}
	//#endregion SpaceShip Right

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&spaceShipLeft.BasicEntity, &spaceShipLeft.RenderComponent, &spaceShipLeft.SpaceComponent)
			sys.Add(&spaceShipRight.BasicEntity, &spaceShipRight.RenderComponent, &spaceShipRight.SpaceComponent)
		case *systems.PlayerControlSystem:
			sys.Add(&spaceShipLeft.BasicEntity, &spaceShipLeft.SpaceComponent, &spaceShipLeft.PlayerId)
			sys.Add(&spaceShipRight.BasicEntity, &spaceShipRight.SpaceComponent, &spaceShipRight.PlayerId)
		case *common.CollisionSystem:
			sys.Add(&spaceShipLeft.BasicEntity, &spaceShipLeft.CollisionComponent, &spaceShipLeft.SpaceComponent)
			sys.Add(&spaceShipRight.BasicEntity, &spaceShipRight.CollisionComponent, &spaceShipRight.SpaceComponent)
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

	fnt := &common.Font{
		URL:  "fonts/PressStart2P-Regular.ttf",
		FG:   color.White,
		Size: 64,
	}

	err = fnt.CreatePreloaded()
	if err != nil {
		panic(err)
	}

	playerOneScore := entity.PlayerScore{BasicEntity: ecs.NewBasic()}
	playerOneScore.RenderComponent.Drawable = common.Text{
		Font: fnt,
		Text: "0",
	}
	playerOneScore.SetShader(common.HUDShader)

	playerOneScore.Position = engo.Point{
		X: 100,
		Y: engo.GameHeight() - spaceShipLeft.Height,
	}
	playerTwoScore := entity.PlayerScore{BasicEntity: ecs.NewBasic()}
	playerTwoScore.RenderComponent.Drawable = common.Text{
		Font: fnt,
		Text: "0",
	}
	playerTwoScore.SetShader(common.HUDShader)

	playerTwoScore.Position = engo.Point{
		X: spaceShipRight.Position.X + 140,
		Y: engo.GameHeight() - spaceShipLeft.Height,
	}
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&playerOneScore.BasicEntity, &playerOneScore.RenderComponent, &playerOneScore.SpaceComponent)
			sys.Add(&playerTwoScore.BasicEntity, &playerTwoScore.RenderComponent, &playerTwoScore.SpaceComponent)
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
