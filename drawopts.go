package main

import "github.com/hajimehoshi/ebiten/v2"

type DrawOption struct {
	opts   *ebiten.DrawImageOptions
	camera *Camera
}

func (do *DrawOption) Translate(x, y float64) {
	do.opts.GeoM.Translate(x, y)
	do.opts.GeoM.Translate(do.camera.X, do.camera.Y)
}
