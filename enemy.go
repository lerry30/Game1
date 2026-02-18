package main

import (
	"fmt"
	"math"

	"game1/animations"
	"game1/constants"
	"game1/spritesheet"
)

type Enemy struct {
	Sprite
	Speed     float64
	Animation map[AnimationState]*animations.Animation
	Target    *Sprite
}

func NewEnemy(filepath string, x, y float64) (*Enemy, error) {
	enemy := Enemy{}

	spriteSheetImgWidth := 64
	spriteSheetImgHeight := 112
	enemySpriteSheet := spritesheet.NewSpriteSheet(spriteSheetImgWidth, spriteSheetImgHeight, constants.TileSize)
	enemyAnimation := map[AnimationState]*animations.Animation{
		Up:    animations.NewAnimation(5, 13, 4, 6.0),
		Down:  animations.NewAnimation(4, 12, 4, 6.0),
		Left:  animations.NewAnimation(6, 14, 4, 6.0),
		Right: animations.NewAnimation(7, 15, 4, 6.0),
	}

	err := enemy.NewSprite(filepath, x, y, enemySpriteSheet, enemyAnimation)
	if err != nil {
		return &Enemy{}, fmt.Errorf("Unable to create enemy at enemy.go %w", err)
	}

	enemy.Speed = 1
	enemy.Animation = enemyAnimation

	return &enemy, nil
}

func (e *Enemy) SetTarget(target *Sprite) {
	e.Target = target
}

func (e *Enemy) MoveX(tileMapWidthPixel float64) {
	e.Dx = 0.0
	distance := e.Distance()
	if distance > 13 && distance < 150 {
		// right
		if e.X < e.Target.X {
			e.Dx = e.Speed
			e.X = math.Min(tileMapWidthPixel-16, e.X+e.Dx)
		}
		// left
		if e.X > e.Target.X {
			e.Dx = -e.Speed
			e.X = math.Max(0.0, e.X+e.Dx)
		}
	}
}

func (e *Enemy) MoveY(tileMapHeightPixel float64) {
	e.Dy = 0.0
	distance := e.Distance()
	if distance > 13 && distance < 150 {
		// up
		if e.Y > e.Target.Y {
			e.Dy = -e.Speed
			e.Y = math.Max(0.0, e.Y+e.Dy)
		}
		// down
		if e.Y < e.Target.Y {
			e.Dy = e.Speed
			e.Y = math.Min(tileMapHeightPixel-16, e.Y+e.Dy)
		}
	}
}

func (e *Enemy) Distance() float64 {
	return math.Hypot(e.X-e.Target.X, e.Y-e.Target.Y)
}
