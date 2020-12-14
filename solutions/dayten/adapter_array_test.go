package dayten

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountAdapterIntervals(t *testing.T) {
	tests := []struct {
		name     string
		adapters []int
		expected int
	}{
		{
			name:     "base",
			adapters: []int{0, 1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19}, // 0 is added by my code
			expected: 35,
		}, {
			name:     "extended",
			adapters: []int{0, 28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3},
			expected: 220,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			sort.Ints(tt.adapters) // I'm lazy and didn't want to retype the input list
			res := CountAdapterIntervals(tt.adapters)
			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestValidCombinations(t *testing.T) {
	tests := []struct {
		name     string
		adapters []int
		expected int
	}{
		{
			name:     "base",
			adapters: []int{0, 1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19}, // 0 is added by my code
			expected: 8,
		}, {
			name:     "extended",
			adapters: []int{0, 28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3},
			expected: 19208,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			sort.Ints(tt.adapters)
			res := DetermineValidCombinations(tt.adapters)
			assert.Equal(t, tt.expected, res)
		})
	}
}
