package main

import (
	"fmt"
	"game1/animations"
	"game1/spritesheet"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type AnimationState uint8

const (
	Down AnimationState = iota
	Up
	Left
	Right
)

type Sprite struct {
	Img             *ebiten.Image
	X, Y, Dx, Dy    float64
	MainSpriteSheet *spritesheet.SpriteSheet
	SpriteAnimation map[AnimationState]*animations.Animation
}

func (s *Sprite) NewSprite(filepath string, x, y float64, ss *spritesheet.SpriteSheet, anim map[AnimationState]*animations.Animation) error {
	spriteImage, _, err := ebitenutil.NewImageFromFile(filepath)
	if err != nil {
		return fmt.Errorf("Opening file error at sprites.go - %w", err)
	}

	s.Img = spriteImage
	s.X = x
	s.Y = y
	s.MainSpriteSheet = ss
	s.SpriteAnimation = anim

	return nil
}

func (s *Sprite) Draw(screen *ebiten.Image, drawOpts *DrawOption) {
	drawOpts.Translate(s.X, s.Y)

	frame := 0
	activeAnim := s.ActiveAnimation()
	if activeAnim != nil {
		frame = activeAnim.GetFrame()
	}

	screen.DrawImage(
		//s.Img.SubImage(image.Rect(0, 0, 16, 16)).(*ebiten.Image),
		s.Img.SubImage(s.MainSpriteSheet.Rect(frame)).(*ebiten.Image),
		drawOpts.opts,
	)

	drawOpts.opts.GeoM.Reset()
}

func (s *Sprite) ActiveAnimation() *animations.Animation {
	if s.Dx < 0.0 {
		return s.SpriteAnimation[Left]
	} else if s.Dx > 0.0 {
		return s.SpriteAnimation[Right]
	} else if s.Dy < 0.0 {
		return s.SpriteAnimation[Up]
	} else if s.Dy > 0.0 {
		return s.SpriteAnimation[Down]
	}

	return nil
}
