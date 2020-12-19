package dayfifteen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterateAndInitFull(t *testing.T) {
	tests := []struct {
		name     string
		starting []int
		count    int
		expected []int
		final    int
	}{
		{
			name:     "base",
			starting: []int{0, 3, 6},
			count:    10,
			expected: []int{0, 3, 3, 1, 0, 4, 0},
			final:    0,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			g := Game{}
			val := g.Initialise(tt.starting)
			for i := 0; i < tt.count-len(tt.starting)-1; i++ {
				assert.Equal(t, tt.expected[i], val)
				val = g.Iterate(val)
			}
			assert.Equal(t, tt.final, val)
		})
	}
}

func TestFull(t *testing.T) {
	tests := []struct {
		name     string
		starting []int
		count    int
		expected int
	}{
		{
			name:     "base",
			starting: []int{0, 3, 6},
			count:    10,
			expected: 0,
		}, {
			name:     "base 2020",
			starting: []int{0, 3, 6},
			count:    2020,
			expected: 436,
		}, {
			name:     "1",
			starting: []int{1, 3, 2},
			count:    2020,
			expected: 1,
		}, {
			name:     "10",
			starting: []int{2, 1, 3},
			count:    2020,
			expected: 10,
		}, {
			name:     "27",
			starting: []int{1, 2, 3},
			count:    2020,
			expected: 27,
		}, {
			name:     "78",
			starting: []int{2, 3, 1},
			count:    2020,
			expected: 78,
		}, {
			name:     "438",
			starting: []int{3, 2, 1},
			count:    2020,
			expected: 438,
		}, {
			name:     "1836",
			starting: []int{3, 1, 2},
			count:    2020,
			expected: 1836,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := Game{}
			val := g.Initialise(tt.starting)
			for i := 0; i < tt.count-len(tt.starting)-1; i++ {
				val = g.Iterate(val)
			}
			assert.Equal(t, tt.expected, val)
		})
	}
}

func TestFullLong(t *testing.T) {
	tests := []struct {
		name     string
		starting []int
		count    int
		expected int
	}{
		{
			name:     "base",
			starting: []int{0, 3, 6},
			count:    30000000,
			expected: 175594,
		}, {
			name:     "1",
			starting: []int{1, 3, 2},
			count:    30000000,
			expected: 2578,
		}, {
			name:     "10",
			starting: []int{2, 1, 3},
			count:    30000000,
			expected: 3544142,
		}, {
			name:     "27",
			starting: []int{1, 2, 3},
			count:    30000000,
			expected: 261214,
		}, {
			name:     "78",
			starting: []int{2, 3, 1},
			count:    30000000,
			expected: 6895259,
		}, {
			name:     "438",
			starting: []int{3, 2, 1},
			count:    30000000,
			expected: 18,
		}, {
			name:     "1836",
			starting: []int{3, 1, 2},
			count:    30000000,
			expected: 362,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := Game{}
			val := g.Initialise(tt.starting)
			for i := 0; i < tt.count-len(tt.starting)-1; i++ {
				val = g.Iterate(val)
			}
			assert.Equal(t, tt.expected, val)
		})
	}
}
