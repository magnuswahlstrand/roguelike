package main

import (
	"golang.org/x/image/colornames"
	"image"
	"image/color"
)

type Entity struct {
	image.Point
	color.Color
	char rune
}

func (e *Entity) Move(dx, dy int) {
	e.X += dx
	e.Y += dy
}

type Player struct {
	*Entity
}

func newPlayer(pos image.Point) *Player {
	p := &Player{
		Entity: &Entity{
			Point: pos,
			Color: colornames.Darkred,
			char:  '@',
		},
	}
	return p
}
