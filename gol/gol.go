// An implementation of Conway's Game of Life.
package gol

// Life stores the state of a round of Conway's Game of Life.
type Life struct {
	w, h int
}

// NewLife returns a new Life game state with a random initial state.
func NewLife(w, h int) *Life {
	return &Life{
		w: w, h: h,
	}
}

func (l *Life) Height() int {
	return l.h
}

func (l *Life) Width() int {
	return l.w
}

func (l *Life) Set(x, y int, b bool) {
	// TODO set specified coords to contain a cell
}

func (l *Life) RuneFor(x, y int) rune {
	// TODO return correct rune
	return ' '
}

// Step advances the game by one instant, recomputing and updating all cells.
func (l *Life) Step() {
	// TODO implement me
}
