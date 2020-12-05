package dayone

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepairReport(t *testing.T) {
	tests := []struct {
		name     string
		input    []int64
		expected int64
	}{
		{
			name:     "base",
			input:    []int64{1721, 979, 366, 299, 675, 1456},
			expected: 514579,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res := RepairReport(tt.input)
			assert.Equal(t, tt.expected, res, "returned value should match expected	")
		})
	}
}

func TestRepairReportExtra(t *testing.T) {
	tests := []struct {
		name     string
		input    []int64
		expected int64
	}{
		{
			name:     "base",
			input:    []int64{1721, 979, 366, 299, 675, 1456},
			expected: 241861950,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res := RepairReportExtra(tt.input)
			assert.Equal(t, tt.expected, res, "returned value should match expected	")
		})
	}
}
