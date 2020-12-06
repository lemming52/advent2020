package dayfive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBoardingPass(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		expectedID int
	}{
		{
			name:       "base",
			input:      "FBFBBFFRLR",
			expectedID: 44*8 + 5,
		}, {
			name:       "base",
			input:      "FBBFFBBLLL",
			expectedID: 51 * 8,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			root := SeatNode{
				rowMin:    0,
				rowMax:    127,
				colMin:    0,
				colMax:    8,
				nodeDepth: 0,
				rowDelta:  64,
				colDelta:  4,
				Children:  0,
				id:        -1,
			}
			res := root.AddSeat(tt.input)
			assert.Equal(t, tt.expectedID, res, "id should match")
		})
	}
}
