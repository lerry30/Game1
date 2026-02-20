package spritesheet

import (
	"image"
)

type SpriteSheet struct {
	WidthInTiles  int
	HeightInTiles int
	TileSize      int
}

func NewSpriteSheet(w, h, t int) *SpriteSheet {
	return &SpriteSheet{
		WidthInTiles:  w,
		HeightInTiles: h,
		TileSize:      t,
	}
}

func (ss *SpriteSheet) Rect(index int) image.Rectangle {
	var x int = index % (ss.WidthInTiles / ss.TileSize)
	var y int = index / (ss.WidthInTiles / ss.TileSize)

	x *= ss.TileSize
	y *= ss.TileSize

	return image.Rect(x, y, x+ss.TileSize, y+ss.TileSize)
}
