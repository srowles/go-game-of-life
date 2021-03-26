package main

import (
	"flag"
	"os"
	"time"

	"github.com/nsf/termbox-go"

	"github.com/srowles/go-game-of-life/gol"
)

const (
	defaultColor = termbox.ColorDefault
	bgColor      = termbox.ColorDefault
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
	termbox.Init()
	termbox.SetOutputMode(termbox.Output256)
	defer termbox.Close()
	go listenToKeyboard()
	ticker := time.NewTicker(time.Millisecond * 150)
	for range ticker.C {
		render(gol)
		gol.Step()
	}
}

func render(gol *gol.Life) {
	termbox.Clear(defaultColor, defaultColor)

	// borders of game
	for y := 1; y < gol.Height()+2; y++ {
		termbox.SetCell(0, y, '║', termbox.ColorBlue, bgColor)
		termbox.SetCell(gol.Width()+2, y, '║', termbox.ColorBlue, bgColor)
	}
	for x := 1; x < gol.Width()+2; x++ {
		termbox.SetCell(x, 0, '═', termbox.ColorBlue, bgColor)
		termbox.SetCell(x, gol.Height()+2, '═', termbox.ColorBlue, bgColor)
	}
	termbox.SetCell(0, 0, '╔', termbox.ColorBlue, bgColor)
	termbox.SetCell(gol.Width()+2, 0, '╗', termbox.ColorBlue, bgColor)
	termbox.SetCell(0, gol.Height()+2, '╚', termbox.ColorBlue, bgColor)
	termbox.SetCell(gol.Width()+2, gol.Height()+2, '╝', termbox.ColorBlue, bgColor)

	// actual game
	for y := 0; y < gol.Height(); y++ {
		for x := 0; x < gol.Width(); x++ {
			b := ' '
			if gol.Get(x, y) {
				b = '*'
			}
			termbox.SetCell(x+1, y+1, b, termbox.ColorGreen, bgColor)
		}
	}

	termbox.Flush()
}

func listenToKeyboard() {
	termbox.SetInputMode(termbox.InputEsc)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				termbox.Close()
				os.Exit(0)
			}
		case termbox.EventError:
			termbox.Close()
			panic(ev.Err)
		}
	}
}
