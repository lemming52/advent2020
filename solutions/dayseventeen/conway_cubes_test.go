package dayseventeen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterate(t *testing.T) {
	turn := 6
	init := [][]bool{
		{false, true, false}, // .#.
		{false, false, true}, // ..#
		{true, true, true},   // ###
	}
	tests := []struct {
		name       string
		dimensions []int
		offset     []int
		expected   int
	}{
		{
			name:       "3",
			dimensions: []int{turn*2 + 1, len(init[0]) + turn*2, len(init[0]) + turn*2},
			offset:     []int{turn, turn, turn},
			expected:   112,
		}, {
			name:       "4",
			dimensions: []int{turn*2 + 1, turn*2 + 1, len(init[0]) + turn*2, len(init[0]) + turn*2},
			offset:     []int{turn, turn, turn, turn},
			expected:   848,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			d := NewDimension(tt.dimensions, 0)
			d.Initalise(init, tt.offset)
			for i := 0; i < turn; i++ {
				d.Iterate()
			}
			assert.Equal(t, tt.expected, d.Total())
		})
	}
}
