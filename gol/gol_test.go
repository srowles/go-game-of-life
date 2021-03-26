package gol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			gol := NewLife(10, 10)
			for k, v := range test.initial {
				gol.Set(k.x, k.y, v)
			}

			gol.Step()
			for k, v := range test.expected {
				assert.Equal(t, v, gol.Get(k.x, k.y))
			}

		})
	}
}
