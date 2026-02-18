package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player    *Player
	enemy     *Enemy
	gMap      *Map
	camera    *Camera
	colliders Colliders
}

func (g *Game) Update() error {
	g.player.ControlX(float64(g.gMap.tileMap.Layers[0].Width) * 16.0)
	g.colliders.CheckCollisionX(&g.player.Sprite)
	g.player.ControlY(float64(g.gMap.tileMap.Layers[0].Height) * 16.0)
	g.colliders.CheckCollisionY(&g.player.Sprite)

	playerAnim := g.player.Sprite.ActiveAnimation()
	if playerAnim != nil {
		playerAnim.Update()
	}

	g.enemy.MoveX(float64(g.gMap.tileMap.Layers[0].Width) * 16.0)
	g.colliders.CheckCollisionX(&g.enemy.Sprite)
	g.enemy.MoveY(float64(g.gMap.tileMap.Layers[0].Height) * 16.0)
	g.colliders.CheckCollisionY(&g.enemy.Sprite)

	enemyAnim := g.enemy.Sprite.ActiveAnimation()
	if enemyAnim != nil {
		enemyAnim.Update()
	}

	screenWidth := 320
	screenHeight := 240

	g.camera.FollowTarget(g.player.X+8, g.player.Y+8, screenWidth, screenHeight)
	g.camera.Constraint(
		float64(g.gMap.tileMap.Layers[0].Width)*16.0,
		float64(g.gMap.tileMap.Layers[0].Height)*16.0,
		screenWidth,
		screenHeight,
	)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{100, 100, 100, 255})

	opts := ebiten.DrawImageOptions{}
	drawOpts := DrawOption{
		opts:   &opts,
		camera: g.camera,
	}

	// Draw Map
	g.gMap.RenderMap(screen, &drawOpts)

	opts.GeoM.Reset()

	// Draw Player
	g.player.Draw(screen, &drawOpts)

	// Draw Enemy
	g.enemy.Draw(screen, &drawOpts)

	// Draw Colliders Debug
	//g.colliders.DrawDebugCollider(screen, g.camera.X, g.camera.Y)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	//return ebiten.WindowSize()
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Game 1")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	// player
	player, err := NewPlayer("assets/images/ninja.png", 100.0, 100.0)
	if err != nil {
		log.Fatal(err)
	}

	// enemy
	enemy, err := NewEnemy("assets/images/skeleton.png", 600.0, 200.0)
	if err != nil {
		log.Fatal(err)
	}
	enemy.SetTarget(&player.Sprite)

	// layers(Map and buildings)
	gMap, err := NewMap("assets/maps/spawn.json")
	if err != nil {
		log.Fatal(err)
	}

	// Camera
	camera := NewCamera(0, 0)

	// Collider
	colliders := NewColliders()

	// Create obstacles
	buildingCoord := gMap.BuildingsCoords()
	for _, coord := range buildingCoord {
		colliders.AddNewCollider(coord)
	}

	game := Game{
		player:    player,
		enemy:     enemy,
		gMap:      gMap,
		camera:    camera,
		colliders: colliders,
	}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
