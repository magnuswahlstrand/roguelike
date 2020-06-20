package main

import (
	"github.com/gdamore/tcell"
	"github.com/kyeett/roguelike/tile"
	"github.com/peterhellberg/gfx"
	"image"
)

func (g *Game) handleEvent() interface{} {
	ev := g.PollEvent()

	switch ev := ev.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyUp:
			return Move{0, -1}
		case tcell.KeyDown:
			return Move{0, 1}
		case tcell.KeyLeft:
			return Move{-1, 0}
		case tcell.KeyRight:
			return Move{1, 0}
		case tcell.KeyEscape, tcell.KeyEnter:
			return Exit{}
		}
	}
	return nil
}

func (g *Game) isBlocked(x int, y int) bool {
	t := g.Grid.Get(image.Pt(x, y))

	switch v := t.(type) {
	case tile.Tile:
		return v.Blocked
	}

	return false
}

func (g *Game) update() error {
	action := g.handleEvent()
	switch v := action.(type) {
	case Move:
		if !g.isBlocked(g.player.X + v.dx, g.player.Y+v.dy) {
			g.player.Move(v.dx, v.dy)
		}


	case Exit:
		return gfx.ErrDone
	}

	return nil
}