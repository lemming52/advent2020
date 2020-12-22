package dayeighteen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
		advanced int
	}{
		{
			name:     "base",
			input:    "1 + 2 * 3 + 4 * 5 + 6",
			expected: 71,
			advanced: 231,
		}, {
			name:     "brackets",
			input:    "1 + (2 * 3) + (4 * (5 + 6))",
			expected: 51,
			advanced: 51,
		}, {
			name:     "26",
			input:    "2 * 3 + (4 * 5)",
			expected: 26,
			advanced: 46,
		}, {
			name:     "437",
			input:    "5 + (8 * 3 + 9 + 3 * 4 * 3)",
			expected: 437,
			advanced: 1445,
		}, {
			name:     "12240",
			input:    "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
			expected: 12240,
			advanced: 669060,
		}, {
			name:     "13632",
			input:    "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			expected: 13632,
			advanced: 23340,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res := Parse(tt.input, false)
			res2 := Parse(tt.input, true)
			assert.Equal(t, tt.expected, res)
			assert.Equal(t, tt.advanced, res2)
		})
	}
}
