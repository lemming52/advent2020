package daynine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
		window   int
	}{
		{
			name: "base",
			input: []int{
				35,
				20,
				15,
				25,
				47,
				40,
				62,
				55,
				65,
				95,
				102,
				117,
				150,
				182,
				127,
				219,
				299,
				277,
				309,
				576,
			},
			expected: 127,
			window:   5,
		}, {
			name: "order",
			input: []int{
				20, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 1, 21, 22, 23, 24, 25, 45, 65,
			},
			expected: 65,
			window:   25,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res := ParseNumbers(tt.input, tt.window)
			assert.Equal(t, tt.expected, res)
		})
	}
}
