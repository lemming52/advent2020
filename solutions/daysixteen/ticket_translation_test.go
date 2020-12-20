package daysixteen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name       string
		conditions [][]string
		entries    string
		expected   int
	}{
		{
			name: "base",
			conditions: [][]string{
				{"1", "3"}, {"5", "7"},
				{"6", "11"}, {"33", "44"},
				{"13", "40"}, {"45", "50"},
			},
			entries:  "7,1,14",
			expected: 0,
		}, {
			name: "base",
			conditions: [][]string{
				{"1", "3"}, {"5", "7"},
				{"6", "11"}, {"33", "44"},
				{"13", "40"}, {"45", "50"},
			},
			entries:  "7,3,47",
			expected: 0,
		}, {
			name: "base",
			conditions: [][]string{
				{"1", "3"}, {"5", "7"},
				{"6", "11"}, {"33", "44"},
				{"13", "40"}, {"45", "50"},
			},
			entries:  "40,4,50",
			expected: 4,
		}, {
			name: "base",
			conditions: [][]string{
				{"1", "3"}, {"5", "7"},
				{"6", "11"}, {"33", "44"},
				{"13", "40"}, {"45", "50"},
			},
			entries:  "55,2,20",
			expected: 55,
		}, {
			name: "base",
			conditions: [][]string{
				{"1", "3"}, {"5", "7"},
				{"6", "11"}, {"33", "44"},
				{"13", "40"}, {"45", "50"},
			},
			entries:  "36,6,12",
			expected: 12,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			v := NewValidator()
			for _, c := range tt.conditions {
				v.AddCondition(c[0], c[1])
			}
			v.Validate(tt.entries)
			assert.Equal(t, tt.expected, v.invalidTotal)
		})
	}
}
