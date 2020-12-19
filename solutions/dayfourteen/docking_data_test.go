package dayfourteen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddValue(t *testing.T) {
	tests := []struct {
		name     string
		value    int
		address  int
		expected int
		mask     string
	}{
		{
			name:     "base",
			value:    11,
			address:  8,
			expected: 73,
			mask:     "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		}, {
			name:     "next",
			value:    101,
			address:  7,
			expected: 101,
			mask:     "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		}, {
			name:     "base",
			value:    0,
			address:  8,
			expected: 64,
			mask:     "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			m := NewMemory()
			m.UpdateMask(tt.mask)
			m.AddValue(tt.address, tt.value)
			assert.Equal(t, tt.expected, m.memory[tt.address])
		})
	}
}

func TestTotal(t *testing.T) {
	mask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"
	inputs := [][]int{
		{8, 11},
		{7, 101},
		{8, 0},
	}
	expected := 165
	m := NewMemory()
	m.UpdateMask(mask)
	for _, i := range inputs {
		m.AddValue(i[0], i[1])
	}
	assert.Equal(t, expected, m.Total())
}
