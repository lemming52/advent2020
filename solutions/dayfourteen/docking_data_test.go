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
		expected uint64
		mask     string
	}{
		{
			name:     "base",
			value:    11,
			address:  8,
			expected: uint64(73),
			mask:     "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		}, {
			name:     "next",
			value:    101,
			address:  7,
			expected: uint64(101),
			mask:     "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		}, {
			name:     "base",
			value:    0,
			address:  8,
			expected: uint64(64),
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
	tests := []struct {
		name     string
		mask     string
		inputs   [][]int
		expected uint64
	}{
		{
			name: "base",
			mask: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
			inputs: [][]int{
				{8, 11},
				{7, 101},
				{8, 0},
			},
			expected: uint64(165),
		}, /*{
			name: "larger",
			mask: "110000011XX0000X101000X10X01XX001011",
			inputs: [][]int{
				{49397, 468472},
				{50029, 23224119},
				{39033, 191252712},
				{37738, 25669},
				{45831, 238647542},
				{55749, 1020},
				{29592, 57996},
			},
			expected: uint64(165),
		},*/
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			m := NewMemory()
			m.UpdateMask(tt.mask)
			for _, i := range tt.inputs {
				m.AddValue(i[0], i[1])
			}
			assert.Equal(t, tt.expected, m.Total())
		})
	}
}
