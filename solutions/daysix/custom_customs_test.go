package daysix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotalQuestions(t *testing.T) {
	tests := []struct {
		name              string
		input             []string
		expectedTotal     int
		expectedAgreement int
	}{
		{
			name:              "base",
			input:             []string{"abc"},
			expectedTotal:     3,
			expectedAgreement: 3,
		}, {
			name:              "split",
			input:             []string{"a", "b", "c"},
			expectedTotal:     3,
			expectedAgreement: 0,
		}, {
			name:              "overlap",
			input:             []string{"ab", "ac"},
			expectedTotal:     3,
			expectedAgreement: 1,
		}, {
			name:              "multiple",
			input:             []string{"a", "a", "a", "a"},
			expectedTotal:     1,
			expectedAgreement: 1,
		}, {
			name:              "single",
			input:             []string{"b"},
			expectedTotal:     1,
			expectedAgreement: 1,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			total, agreement := TotalQuestions(tt.input)
			assert.Equal(t, tt.expectedTotal, total, "number of answered questions should match")
			assert.Equal(t, tt.expectedAgreement, agreement, "number of agreed questions should match")
		})
	}
}
