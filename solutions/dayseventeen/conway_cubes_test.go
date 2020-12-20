package dayseventeen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterate(t *testing.T) {
	init := [][]bool{
		{false, true, false}, // .#.
		{false, false, true}, // ..#
		{true, true, true},   // ###
	}
	turn := 6
	x, y, z := len(init[0])+turn*2, len(init[0])+turn*2, turn*2+1
	offsetX, offsetY := turn, turn
	expected := 112
	c := NewCube(3, z, y, x)
	c.Initalise(init, offsetX, offsetY)
	for i := 0; i < turn; i++ {
		c.Iterate()
	}
	assert.Equal(t, expected, c.Total())
}
