package main

import (
	"github.com/gdamore/tcell"
	"github.com/kyeett/roguelike/tile"
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

			t := g.Grid.Get(image.Pt(x, y)).(*tile.Tile)
			if t.Blocked {
				g.SetCell(x, y, darkWall, ' ')
			} else {
				g.SetCell(x, y, darkGround, ' ')
			}
		}
	}
	// Draw all entities in the list
	for _, e := range g.entities {
		re, gr, bl, _ := e.Color.RGBA()
		g.SetCell(e.X, e.Y, tcell.StyleDefault.Foreground(tcell.NewRGBColor(int32(re), int32(gr), int32(bl))), e.char)
	}
	g.Show()

	g.Clear()
}
