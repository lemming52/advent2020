package dayfourteen

import (
	"log"
	"reflect"
	"regexp"
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
		},
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

func TestMemory2AddValue(t *testing.T) {
	tests := []struct {
		name        string
		value       int
		address     int
		expected    uint64
		expectedMap map[uint64]uint64
		mask        string
	}{
		{
			name:     "base",
			value:    100,
			address:  42,
			expected: 400,
			expectedMap: map[uint64]uint64{
				26: 100,
				27: 100,
				58: 100,
				59: 100,
			},
			mask: "000000000000000000000000000000X1001X",
		}, {
			name:     "next",
			value:    1,
			address:  26,
			expected: 8,
			expectedMap: map[uint64]uint64{
				16: 1,
				17: 1,
				18: 1,
				19: 1,
				24: 1,
				25: 1,
				26: 1,
				27: 1,
			},
			mask: "00000000000000000000000000000000X0XX",
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			m := NewMemory2()
			m.UpdateMask(tt.mask)
			m.AddValue(tt.address, tt.value)
			assert.Equal(t, tt.expected, m.Total())
			if !reflect.DeepEqual(tt.expectedMap, m.memory) {
				t.Errorf("Memory should match %v %v", tt.expectedMap, m.memory)
			}
		})
	}
}

func TestTotal2(t *testing.T) {
	tests := []struct {
		name     string
		mask     string
		lines    []string
		expected uint64
	}{
		{
			name: "base",
			lines: []string{
				"mask = 000000000000000000000000000000X1001X",
				"mem[42] = 100",
				"mask = 00000000000000000000000000000000X0XX",
				"mem[26] = 1",
			},
			expected: uint64(208),
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			m := NewMemory2()
			maskPattern, err := regexp.Compile(`mask = ([0,1,X]{36})`)
			if err != nil {
				log.Fatal(err)
			}
			valPattern, err := regexp.Compile(`mem\[(\d+)\] = (\d+)`)
			if err != nil {
				log.Fatal(err)
			}
			for _, l := range tt.lines {
				parseLine(l, maskPattern, valPattern, m)
			}
			assert.Equal(t, tt.expected, m.Total())
		})
	}
}
