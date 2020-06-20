package main

import (
	"github.com/gdamore/tcell"
	"github.com/peterhellberg/gfx"
	"image"
	"log"
)

type Game struct {
	tcell.Screen
	player *Player
}

func newGame() *Game {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Init(); err != nil {
		log.Fatal(err)
	}

	s.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorGhostWhite).
		Background(tcell.ColorPurple))

	pos := image.Pt(1, 1)
	p := newPlayer(pos)

	return &Game{s, p}
}

type Move struct {
	dx, dy int
}

type Exit struct {}

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

func (g *Game) update() error {
	action := g.handleEvent()
	switch v := action.(type) {
	case Move:
		g.player.Move(v.dx, v.dy)
	case Exit:
		return gfx.ErrDone
	}

	return nil
}

func (g *Game) draw() {
	g.Clear()
	g.SetCell(g.player.X, g.player.Y, tcell.StyleDefault, g.player.char)
	g.Screen.Sync()
}
