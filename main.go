package main

import (
	"flag"

	"github.com/srowles/go-game-of-life/gol"
	"github.com/srowles/go-game-of-life/ui/terminal"
)

func main() {
	var width, height int
	flag.IntVar(&width, "width", 60, "width of the board")
	flag.IntVar(&height, "height", 20, "height of the board")
	flag.Parse()

	gol := gol.NewLife(width, height)

	// a glider
	gol.Set(10, 10, true)
	gol.Set(11, 10, true)
	gol.Set(12, 10, true)
	gol.Set(10, 11, true)
	gol.Set(11, 12, true)

	// static square
	gol.Set(2, 2, true)
	gol.Set(2, 3, true)
	gol.Set(3, 2, true)
	gol.Set(3, 3, true)

	// becomes a "circle"
	gol.Set(30, 10, true)
	gol.Set(31, 10, true)
	gol.Set(32, 10, true)
	gol.Set(30, 11, true)

	//
	gol.Set(50, 8, true)
	gol.Set(51, 8, true)
	gol.Set(53, 8, true)
	gol.Set(54, 8, true)
	gol.Set(51, 9, true)
	gol.Set(53, 9, true)
	gol.Set(51, 10, true)
	gol.Set(53, 10, true)
	gol.Set(52, 11, true)

	ui := terminal.New(gol)
	ui.Run()
}
