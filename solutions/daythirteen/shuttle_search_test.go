package daythirteen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindEarliest(t *testing.T) {
	ids := []int{7, 13, 59, 31, 19}
	departureTime := 939
	expected := 295
	res := FindEarliest(ids, departureTime)
	assert.Equal(t, expected, res)
}

func TestFindMagicalDepartureTime(t *testing.T) {
	tests := []struct {
		name     string
		shuttles []int
		expected int
	}{
		{
			name:     "base",
			shuttles: []int{17, -1, 13, 19},
			expected: 3417,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res := FindMagicalDepartureTime(tt.shuttles)
			assert.Equal(t, tt.expected, res)
		})
	}
}
