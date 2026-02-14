package main

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	Img          *ebiten.Image
	X, Y, Dx, Dy float64
}

func (s *Sprite) NewSprite(filepath string, x, y float64) error {
	spriteImage, _, err := ebitenutil.NewImageFromFile(filepath)
	if err != nil {
		return fmt.Errorf("Opening file error at sprites.go - %w", err)
	}

	s.Img = spriteImage
	s.X = x
	s.Y = y

	return nil
}

func (s *Sprite) Draw(screen *ebiten.Image, drawOpts *DrawOption) {
	drawOpts.Translate(s.X, s.Y)

	screen.DrawImage(
		s.Img.SubImage(image.Rect(0, 0, 16, 16)).(*ebiten.Image),
		drawOpts.opts,
	)

	drawOpts.opts.GeoM.Reset()
}
