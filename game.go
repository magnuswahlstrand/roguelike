package main

import (
	"github.com/gdamore/tcell"
	"github.com/kyeett/collections/grid"
	"github.com/kyeett/roguelike/tile"
	"image"
	"log"
)

type Game struct {
	tcell.Screen
	player   *Player
	entities []*Entity

	*grid.Grid
}

const (
	mapWidth  = 60
	mapHeight = 40
)

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

	pos := image.Pt(mapWidth/2, mapHeight/2)
	p := newPlayer(pos)

	entities := []*Entity{
		p.Entity,
	}

	gr := grid.New(mapWidth, mapHeight)
	t := tile.Tile{true, true}
	gr.Set(image.Pt(30, 22), t)
	gr.Set(image.Pt(31, 22), t)
	gr.Set(image.Pt(32, 22), t)

	return &Game{s, p, entities, gr}
}

type Move struct {
	dx, dy int
}

type Exit struct{}

