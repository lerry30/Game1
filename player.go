package main

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Sprite
	Speed float64
}

func NewPlayer(filepath string, x, y float64) (*Player, error) {
	player := Player{}
	err := player.NewSprite(filepath, x, y)
	if err != nil {
		return &Player{}, fmt.Errorf("Unable to create player at player.go %w", err)
	}

	player.Speed = 2

	return &player, nil
}

func (p *Player) ControlX(tileMapWidthPixel float64) {
	p.Dx = 0.0
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.Dx = p.Speed
		p.X = math.Min(tileMapWidthPixel-16, p.X+p.Dx)
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.Dx = -p.Speed
		p.X = math.Max(0.0, p.X+p.Dx)
	}
}

func (p *Player) ControlY(tileMapHeightPixel float64) {
	p.Dy = 0.0
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.Dy = -p.Speed
		p.Y = math.Max(0.0, p.Y+p.Dy)
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.Dy = p.Speed
		p.Y = math.Min(tileMapHeightPixel-16, p.Y+p.Dy)
	}
}
