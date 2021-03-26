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
	flag.BoolVar(&glider, "glider", false, "start wit a single glider")
	flag.Parse()

	gol := gol.NewLife(60, 20)
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
	ticker := time.NewTicker(time.Millisecond * 200)
	for range ticker.C {
		render(gol)
		gol.Step()
	}
}

func render(gol *gol.Life) {
	termbox.Clear(defaultColor, defaultColor)

	for y := 0; y < gol.Height(); y++ {
		for x := 0; x < gol.Width(); x++ {
			b := gol.RuneFor(x, y)
			termbox.SetCell(x, y, b, termbox.ColorGreen, bgColor)
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
