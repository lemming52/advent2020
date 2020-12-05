package daythree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeCounter(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
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
			expected: 7,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res := TreeCounter(tt.input)
			assert.Equal(t, tt.expected, res, "number of encountered trees should match")
		})
	}
}
