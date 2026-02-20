package main

import (
	"math"
)

type Camera struct {
	X, Y float64
}

func NewCamera(x, y float64) *Camera {
	return &Camera{
		X: x,
		Y: y,
	}
}

func (c *Camera) FollowTarget(targetX, targetY float64, screenWidth, screenHeight int) {
	c.X = -targetX + (float64(screenWidth) / 2)
	c.Y = -targetY + (float64(screenHeight) / 2)
}

func (c *Camera) Constraint(tileMapWidthPixel, tileMapHeightPixel float64, screenWidth, screenHeight int) {
	c.X = math.Min(c.X, 0.0)
	c.Y = math.Min(c.Y, 0.0)

	c.X = math.Max(c.X, float64(screenWidth)-tileMapWidthPixel)
	c.Y = math.Max(c.Y, float64(screenHeight)-tileMapHeightPixel)
}
