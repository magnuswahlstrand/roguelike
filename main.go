package main

import (
	"github.com/gdamore/tcell"
	"time"
)


func main() {

	g := newGame()
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)

	quit := make(chan struct{})
	go func() {
		for {

			g.draw()
			err := g.update()
			if err != nil {
				close(quit)
				return
			}
		}
	}()

loop:
	for {
		select {
		case <-quit:
			break loop
		case <-time.After(time.Millisecond * 50):
		}
	}

	g.Fini()
}

