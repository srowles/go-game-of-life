package terminal

import (
	"fmt"
	"os"
	"time"

	"github.com/nsf/termbox-go"
	"github.com/srowles/go-game-of-life/ui"
)

const (
	defaultColor = termbox.ColorDefault
	bgColor      = termbox.ColorDefault
)

// Terminal holds state for terminal based rendering
type Terminal struct {
	game       ui.GameOfLife
	generation int
}

// New creates a new terminal UI game of life
func New(game ui.GameOfLife) *Terminal {
	return &Terminal{
		game: game,
	}
}

// Run starts the rendering and steps the game
func (t *Terminal) Run() {

	termbox.Init()
	termbox.SetOutputMode(termbox.Output256)
	defer termbox.Close()
	go listenToKeyboard()

	t.render()
	// Pause to appreciate initial state
	<-time.After(2 * time.Second)
	ticker := time.NewTicker(time.Millisecond * 150)
	for range ticker.C {
		t.render()
		t.game.Step()
	}
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

func (t *Terminal) render() {
	termbox.Clear(defaultColor, defaultColor)

	// borders of game
	for y := 1; y < t.game.Height()+2; y++ {
		termbox.SetCell(0, y, '║', termbox.ColorBlue, bgColor)
		termbox.SetCell(t.game.Width()+2, y, '║', termbox.ColorBlue, bgColor)
	}
	for x := 1; x < t.game.Width()+2; x++ {
		termbox.SetCell(x, 0, '═', termbox.ColorBlue, bgColor)
		termbox.SetCell(x, t.game.Height()+2, '═', termbox.ColorBlue, bgColor)
	}
	termbox.SetCell(0, 0, '╔', termbox.ColorBlue, bgColor)
	termbox.SetCell(t.game.Width()+2, 0, '╗', termbox.ColorBlue, bgColor)
	termbox.SetCell(0, t.game.Height()+2, '╚', termbox.ColorBlue, bgColor)
	termbox.SetCell(t.game.Width()+2, t.game.Height()+2, '╝', termbox.ColorBlue, bgColor)

	// actual game
	for y := 0; y < t.game.Height(); y++ {
		for x := 0; x < t.game.Width(); x++ {
			b := ' '
			if t.game.Get(x, y) {
				b = '█'
			}
			termbox.SetCell(x+1, y+1, b, termbox.ColorGreen, bgColor)
		}
	}
	t.displayGeneration()

	termbox.Flush()
	t.generation++
}

func (t *Terminal) displayGeneration() {
	output := fmt.Sprintf("Generation: %d     ", t.generation)
	for x, r := range output {
		termbox.SetCell(x, t.game.Height()+3, r, termbox.ColorGreen, bgColor)
	}
}
