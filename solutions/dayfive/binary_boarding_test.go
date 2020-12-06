package dayfive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBoardingPass(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedRow    int
		expectedColumn int
	}{
		{
			name:           "base",
			input:          "FBFBBFFRLR",
			expectedRow:    44,
			expectedColumn: 5,
		}, {
			name:           "base",
			input:          "FBBFFBBLLL",
			expectedRow:    51,
			expectedColumn: 0,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			row, col := ParseBoardingPass(tt.input)
			assert.Equal(t, tt.expectedRow, row, "row should match")
			assert.Equal(t, tt.expectedColumn, col, "column should match")
		})
	}
}
