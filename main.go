package main

import (
	"flag"

	"github.com/srowles/go-game-of-life/gol"
	"github.com/srowles/go-game-of-life/ui/terminal"
)

func main() {
	var glider bool
	var width, height int
	flag.BoolVar(&glider, "glider", false, "start wit a single glider")
	flag.IntVar(&width, "width", 60, "width of the board")
	flag.IntVar(&height, "height", 20, "height of the board")
	flag.Parse()

	gol := gol.NewLife(width, height)
	if glider {
		gol.Set(10, 10, true)
		gol.Set(11, 10, true)
		gol.Set(12, 10, true)
		gol.Set(10, 11, true)
		gol.Set(11, 12, true)
	}

	ui := terminal.New(gol)
	ui.Run()
}
