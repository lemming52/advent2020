package dayeight

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	tests := []struct {
		name          string
		index         int
		acc           int
		instruction   string
		expectedIndex int
		expectedAcc   int
	}{
		{
			name:          "base",
			index:         0,
			acc:           0,
			instruction:   "nop +0",
			expectedIndex: 1,
			expectedAcc:   0,
		}, {
			name:          "acc +",
			index:         0,
			acc:           0,
			instruction:   "acc +5",
			expectedIndex: 1,
			expectedAcc:   5,
		}, {
			name:          "acc -",
			index:         0,
			acc:           10,
			instruction:   "acc -5",
			expectedIndex: 1,
			expectedAcc:   5,
		}, {
			name:          "jmp +",
			index:         0,
			acc:           0,
			instruction:   "jmp +5",
			expectedIndex: 5,
			expectedAcc:   0,
		}, {
			name:          "jmp -",
			index:         10,
			acc:           0,
			instruction:   "jmp -5",
			expectedIndex: 5,
			expectedAcc:   0,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			processor := Processor{
				acc:   tt.acc,
				index: tt.index,
			}
			processor.Execute(tt.instruction)
			assert.Equal(t, tt.expectedAcc, processor.acc)
			assert.Equal(t, tt.expectedIndex, processor.index)
		})
	}
}

func TestExecuteInstructions(t *testing.T) {
	tests := []struct {
		name         string
		instructions []string
		expected     int
	}{
		{
			name: "base",
			instructions: []string{
				"nop +0",
				"acc +1",
				"jmp +4",
				"acc +3",
				"jmp -3",
				"acc -99",
				"acc +1",
				"jmp -4",
				"acc +6",
			},
			expected: 5,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			processor := Processor{
				executed: map[int]bool{},
			}
			res, success := processor.ExecuteInstructions(tt.instructions, -1, "")
			assert.Equal(t, tt.expected, res)
			assert.Equal(t, false, success)
		})
	}
}

func TestExecuteInstructionsSuccess(t *testing.T) {
	tests := []struct {
		name         string
		instructions []string
		expected     int
	}{
		{
			name: "base",
			instructions: []string{
				"nop +0",
				"acc +1",
				"jmp +4",
				"acc +3",
				"jmp -3",
				"acc -99",
				"acc +1",
				"jmp -4",
				"acc +6",
			},
			expected: 8,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			processor := Processor{
				executed: map[int]bool{},
			}
			res := processor.ExecuteInstructionsSuccess(tt.instructions)
			assert.Equal(t, tt.expected, res)
		})
	}
}
