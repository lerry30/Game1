package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"game1/constants"
)

type Rect struct {
	X, Y, Width, Height int
}

type Colliders []image.Rectangle

func NewColliders() Colliders {
	return make(Colliders, 0)
}

func (c *Colliders) AddNewCollider(rect *Rect) {
	*c = append(*c, image.Rect(
		rect.X,
		rect.Y,
		rect.X+rect.Width,
		rect.Y+rect.Height,
	))
}

func (c *Colliders) DrawDebugCollider(screen *ebiten.Image, camX, camY float64) {
	for _, collider := range *c {
		vector.StrokeRect(screen,
			float32(collider.Min.X)+float32(camX),
			float32(collider.Min.Y)+float32(camY),
			float32(collider.Dx()),
			float32(collider.Dy()),
			1.0,
			color.RGBA{255, 0, 0, 255},
			true,
		)
	}
}

func (c *Colliders) CheckCollisionX(sprite *Sprite) {
	for _, collider := range *c {
		if collider.Overlaps(
			image.Rect(
				int(sprite.X),
				int(sprite.Y),
				int(sprite.X+constants.TileSize),
				int(sprite.Y+constants.TileSize))) {
			if sprite.Dx < 0.0 { // left
				sprite.X = float64(collider.Min.X + collider.Dx())
			} else if sprite.Dx > 0.0 { // right
				sprite.X = float64(collider.Min.X) - 16.0
			}
		}
	}
}

func (c *Colliders) CheckCollisionY(sprite *Sprite) {
	for _, collider := range *c {
		if collider.Overlaps(
			image.Rect(
				int(sprite.X),
				int(sprite.Y),
				int(sprite.X+constants.TileSize),
				int(sprite.Y+constants.TileSize))) {
			if sprite.Dy < 0.0 { // up
				sprite.Y = float64(collider.Min.Y + collider.Dy())
			} else if sprite.Dy > 0.0 { // down
				sprite.Y = float64(collider.Min.Y) - 16.0
			}

		}
	}
}
