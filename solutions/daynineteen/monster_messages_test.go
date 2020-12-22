package daynineteen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRules(t *testing.T) {
	c := &Codex{}
	c.rules = map[int]*Rule{
		0: &Rule{
			name:     0,
			ruleType: SEQUENCE,
			codex:    c,
			sequence: []int{4, 1, 5},
		},
		1: &Rule{
			name:      1,
			ruleType:  ORSEQUENCE,
			codex:     c,
			sequence:  []int{2, 3},
			secondary: []int{3, 2},
		},
		2: &Rule{
			name:      2,
			ruleType:  ORSEQUENCE,
			codex:     c,
			sequence:  []int{4, 4},
			secondary: []int{5, 5},
		},
		3: &Rule{
			name:      3,
			ruleType:  ORSEQUENCE,
			codex:     c,
			sequence:  []int{4, 5},
			secondary: []int{5, 4},
		},
		4: &Rule{
			name:     4,
			ruleType: CHARACTER,
			codex:    c,
			value:    'a',
		},
		5: &Rule{
			name:     5,
			ruleType: CHARACTER,
			codex:    c,
			value:    'b',
		},
	}
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "base",
			input:    "ababbb",
			expected: true,
		}, {
			name:     "a",
			input:    "bababa",
			expected: false,
		}, {
			name:     "b",
			input:    "abbbab",
			expected: true,
		}, {
			name:     "c",
			input:    "aaabbb",
			expected: false,
		}, {
			name:     "d",
			input:    "aaaabbb",
			expected: false,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res, _ := c.rules[0].Apply(tt.input)

			assert.Equal(t, tt.expected, res)
		})
	}
}
