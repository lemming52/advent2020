package daysix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotalQuestions(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name:     "base",
			input:    []string{"abc"},
			expected: 3,
		}, {
			name:     "split",
			input:    []string{"a", "b", "c"},
			expected: 3,
		}, {
			name:     "overlap",
			input:    []string{"ab", "ac"},
			expected: 3,
		}, {
			name:     "multiple",
			input:    []string{"a", "a", "a", "a"},
			expected: 1,
		}, {
			name:     "single",
			input:    []string{"b"},
			expected: 1,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res := TotalQuestions(tt.input)
			assert.Equal(t, tt.expected, res, "number of answered questions should match")
		})
	}
}
