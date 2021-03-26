package gol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Simple example test case which can be used for testing RuleOne
func TestRuleOne(t *testing.T) {
	type point struct {
		x, y int
	}
	tests := map[string]struct {
		initial  map[point]bool
		expected map[point]bool
	}{
		"empty": {},
		// TODO add test cases here to support rule 1
		// Don't forget to test edge cases (and edges of the board!)
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			gol := NewLife(10, 10)

			// Initialise the board
			for k, v := range test.initial {
				gol.Set(k.x, k.y, v)
			}

			// Run one single step
			gol.Step()

			// Check that every square in the board
			// is now populated as expected, making use
			// of the default 'false' behaviour of maps.
			for y := 0; y < gol.Width(); y++ {
				for x := 0; x < gol.Height(); x++ {
					assert.Equal(t, test.expected[point{x: x, y: y}], gol.Get(x, y))
				}
			}
		})
	}
}
