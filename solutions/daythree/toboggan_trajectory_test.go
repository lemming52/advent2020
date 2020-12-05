package daythree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeCounter(t *testing.T) {
	tests := []struct {
		name            string
		input           []string
		widthIncrement  int
		heightIncrement int
		expected        int
	}{
		{
			name: "base",
			input: []string{
				"..##.......",
				"#...#...#..",
				".#....#..#.",
				"..#.#...#.#",
				".#...##..#.",
				"..#.##.....",
				".#.#.#....#",
				".#........#",
				"#.##...#...",
				"#...##....#",
				".#..#...#.#",
			},
			widthIncrement:  3,
			heightIncrement: 1,
			expected:        7,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res := TreeCounter(tt.input, tt.widthIncrement, tt.heightIncrement)
			assert.Equal(t, tt.expected, res, "number of encountered trees should match")
		})
	}
}
