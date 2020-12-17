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
		}, {
			name:     "2",
			shuttles: []int{67, 7, 59, 61},
			expected: 754018,
		}, {
			name:     "3",
			shuttles: []int{67, -1, 7, 59, 61},
			expected: 779210,
		}, {
			name:     "4",
			shuttles: []int{67, 7, -1, 59, 61},
			expected: 1261476,
		}, {
			name:     "5",
			shuttles: []int{1789, 37, 47, 1889},
			expected: 1202161486,
		}, {
			name:     "full",
			shuttles: []int{7, 13, -1, -1, 59, -1, 31, 19},
			expected: 1068781,
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
