package main

import (
	"github.com/gdamore/tcell"
	"github.com/kyeett/collections/grid"
	"github.com/kyeett/roguelike/room"
	"github.com/kyeett/roguelike/tile"
	"github.com/peterhellberg/gfx"
	"golang.org/x/image/colornames"
	"image"
	"log"
	"math/rand"
)

type Game struct {
	tcell.Screen
	player   *Player
	entities []*Entity

	*grid.Grid
}

const (
	mapWidth  = 80
	mapHeight = 50

	roomMinSize = 6
	roomMaxSize = 10
	maxRooms    = 30
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
		&Entity{
			Point: p.Point.Add(image.Point{-5, 0}),
			Color: colornames.Yellow,
			char:  '@',
		},
	}

	gr := grid.New(mapWidth, mapHeight)
	setupTiles(gr)

	// Make map
	var rooms []room.Room
	for i := 0; i < maxRooms; i++ {

		w := rand.Intn(roomMaxSize-roomMinSize+1) + roomMinSize
		h := rand.Intn(roomMaxSize-roomMinSize+1) + roomMinSize

		x := rand.Intn(mapWidth - w - 1)
		y := rand.Intn(mapHeight - h - 1)

		r := room.New(image.Pt(x, y), w, h)

		// Skip room, if it intersects other
		if intersectsOtherRoom(rooms, r) {
			continue
		}

		// Paint room
		setupRoom(gr, r)

		newX, newY := r.Center()
		if i == 0 {

			p.X = newX
			p.Y = newY
		} else {
			// Connect it to the previous room
			prevX, prevY := rooms[len(rooms)-1].Center()

			switch rand.Intn(2) {
			case 0:
				setupHTunnel(gr, prevX, newX, prevY)
				setupVTunnel(gr, prevY, newY, newX)

			case 1:
				setupVTunnel(gr, prevY, newY, prevX)
				setupHTunnel(gr, prevX, newX, newY)
				p.X = newX
				p.Y = newY
			}

		}

		rooms = append(rooms, r)
	}

	return &Game{s, p, entities, gr}
}

func intersectsOtherRoom(rooms []room.Room, r room.Room) bool {
	for _, other := range rooms {
		if r.Intersects(other) {
			return true
		}
	}
	return false
}

func setupHTunnel(gr *grid.Grid, x1 int, x2 int, y int) {
	for x := gfx.IntMin(x1, x2); x <= gfx.IntMax(x1, x2); x++ {
		gr.Get(image.Pt(x, y)).(*tile.Tile).Blocked = false
		gr.Get(image.Pt(x, y)).(*tile.Tile).BlockSight = false
	}
}

func setupVTunnel(gr *grid.Grid, y1, y2, x int) {
	for y := gfx.IntMin(y1, y2); y <= gfx.IntMax(y1, y2); y++ {
		gr.Get(image.Pt(x, y)).(*tile.Tile).Blocked = false
		gr.Get(image.Pt(x, y)).(*tile.Tile).BlockSight = false
	}
}

func setupTiles(gr *grid.Grid) {
	for y := 0; y < gr.Rows(); y++ {
		for x := 0; x < gr.Cols(); x++ {
			gr.Set(image.Pt(x, y), &tile.Tile{true, true})
		}
	}
}

func setupRoom(gr *grid.Grid, r room.Room) {
	for y := r.Min.Y + 1; y < r.Max.Y; y++ {
		for x := r.Min.X + 1; x < r.Max.X; x++ {
			gr.Get(image.Pt(x, y)).(*tile.Tile).Blocked = false
			gr.Get(image.Pt(x, y)).(*tile.Tile).BlockSight = false
		}
	}
}

type Move struct {
	dx, dy int
}

type Exit struct{}
