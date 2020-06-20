package main

import (
	"github.com/gdamore/tcell"
	"image"
)

var darkWall = tcell.StyleDefault.
	Background(tcell.NewRGBColor(0, 0, 100))

var darkGround = tcell.StyleDefault.
	Background(tcell.NewRGBColor(50, 50, 150))

func (g *Game) draw() {

	// Draw all the tiles in the game map
	for y := 0; y < mapHeight; y++ {
		for x := 0; x < mapWidth; x++ {

			tile := g.Grid.Get(image.Pt(x, y))
			isWall := tile != nil
			if isWall {
				g.SetCell(x, y, darkWall, ' ')
			} else {
				g.SetCell(x, y, darkGround, ' ')
			}
		}
	}
	// Draw all entities in the list
	for _, e := range g.entities {
		g.SetCell(e.X, e.Y, tcell.StyleDefault, e.char)
	}
	g.Show()

	g.Clear()
}
