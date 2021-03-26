// An implementation of Conway's Game of Life.
package gol

// Life stores the state of a round of Conway's Game of Life.
type Life struct {
	width  int
	height int
}

// NewLife returns a new Life game state with a random initial state.
func NewLife(width int, height int) *Life {
	return &Life{
		width:  width,
		height: height,
	}
}

func (l *Life) Height() int {
	return l.height
}

func (l *Life) Width() int {
	return l.width
}

func (l *Life) Set(x, y int, alive bool) {
	// TODO implement me
}

func (l *Life) Get(x, y int) bool {
	// TODO implement me
	return false
}

// Step advances the game by one instant, recomputing and updating all cells.
func (l *Life) Step() {
	// TODO implement me
}
