package daysixteen

import (
	"reflect"
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
			v.AddTicket(tt.entries)
			for _, c := range tt.conditions {
				v.AddCondition("", c[0], c[1])
			}
			v.Validate(tt.entries)
			assert.Equal(t, tt.expected, v.invalidTotal)
		})
	}
}

func TestAddCondition(t *testing.T) {
	conditions := [][]string{
		{"class", "1", "3"}, {"title", "5", "7"}, {"extra", "3", "4"},
	}
	expected := map[int][]string{
		1: []string{"class"},
		2: []string{"class"},
		3: []string{"class", "extra"},
		4: []string{"extra"},
		5: []string{"title"},
		6: []string{"title"},
		7: []string{"title"},
	}
	v := NewValidator()
	for _, c := range conditions {
		v.AddCondition(c[0], c[1], c[2])
	}
	if !reflect.DeepEqual(expected, v.values) {
		t.Errorf("valid should match %v %v", expected, v.values)
	}
}

func TestDetermineFields(t *testing.T) {
	conditions := [][]string{
		{"class", "0", "1"}, {"class", "4", "19"},
		{"row", "0", "5"}, {"row", "8", "19"},
		{"seat", "0", "13"}, {"seat", "16", "19"},
	}
	ticket := "11,12,13"
	rows := []string{
		"3,9,18",
		"15,1,5",
		"5,14,9",
	}
	expectedValues := map[int][]string{
		0:  []string{"class", "row", "seat"},
		1:  []string{"class", "row", "seat"},
		2:  []string{"row", "seat"},
		3:  []string{"row", "seat"},
		4:  []string{"class", "row", "seat"},
		5:  []string{"class", "row", "seat"},
		6:  []string{"class", "seat"},
		7:  []string{"class", "seat"},
		8:  []string{"class", "row", "seat"},
		9:  []string{"class", "row", "seat"},
		10: []string{"class", "row", "seat"},
		11: []string{"class", "row", "seat"},
		12: []string{"class", "row", "seat"},
		13: []string{"class", "row", "seat"},
		14: []string{"class", "row"},
		15: []string{"class", "row"},
		16: []string{"class", "row", "seat"},
		17: []string{"class", "row", "seat"},
		18: []string{"class", "row", "seat"},
		19: []string{"class", "row", "seat"},
	}
	expectedValid := [][]int{
		[]int{3, 15, 5},
		[]int{9, 1, 14},
		[]int{18, 5, 9},
	}
	expectedFields := map[int][]string{
		0: []string{"row"},
		1: []string{"class", "row"},
		2: []string{"class", "row", "seat"},
	}
	v := NewValidator()
	v.AddTicket(ticket)
	for _, c := range conditions {
		v.AddCondition(c[0], c[1], c[2])
	}
	for _, r := range rows {
		v.Validate(r)
	}
	fields := v.DetermineFields()
	if !reflect.DeepEqual(expectedValues, v.values) {
		t.Errorf("values should match %v %v", expectedValues, v.values)
	}
	if !reflect.DeepEqual(expectedValid, v.valid) {
		t.Errorf("valid should match %v %v", expectedValid, v.valid)
	}
	for i, f := range fields {
		if !reflect.DeepEqual(expectedFields[i], f) {
			t.Errorf("fields should match %v %v", expectedFields[i], f)
		}
	}
}
