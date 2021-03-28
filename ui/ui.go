package ui

// GameOfLife defines the methods requried to render a game of life
type GameOfLife interface {
	// Height returns the height of the game area
	Height() int

	// Width returns the width of the game area
	Width() int

	// Set allows initial data to be set
	Set(x, y int, alive bool)

	// Get allows the UI to get the value at a location for rendering
	Get(x, y int) bool

	// Step advances the game by one instant, recomputing and updating all cells.
	Step()
}
