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
	}{
		{
			name:     "base",
			input:    "1 + 2 * 3 + 4 * 5 + 6",
			expected: 71,
		}, {
			name:     "brackets",
			input:    "1 + (2 * 3) + (4 * (5 + 6))",
			expected: 51,
		}, {
			name:     "26",
			input:    "2 * 3 + (4 * 5)",
			expected: 26,
		}, {
			name:     "437",
			input:    "5 + (8 * 3 + 9 + 3 * 4 * 3)",
			expected: 437,
		}, {
			name:     "12240",
			input:    "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
			expected: 12240,
		}, {
			name:     "13632",
			input:    "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			expected: 13632,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res := Parse(tt.input)
			assert.Equal(t, tt.expected, res)
		})
	}
}
