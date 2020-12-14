package dayeleven

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddRow(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectedRow []rune
	}{
		{
			name:        "base",
			input:       "L.LL.LL.LL",
			expectedRow: []rune{'L', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			seating := NewSeating()
			seating.AddRow(tt.input)
			if !reflect.DeepEqual(seating.seats[0], tt.expectedRow) {
				t.Errorf("Parse Rows should match")
			}
			assert.Equal(t, 1, seating.height)
		})
	}
}

func TestCountOccupied(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name: "empty",
			input: []string{
				"L.LL.LL.LL",
				"LLLLLLL.LL",
				"L.L.L..L..",
				"LLLL.LL.LL",
				"L.LL.LL.LL",
				"L.LLLLL.LL",
				"..L.L.....",
				"LLLLLLLLLL",
				"L.LLLLLL.L",
				"L.LLLLL.LL",
			},
			expected: 0,
		}, {
			name: "full",
			input: []string{
				"#.##.##.##",
				"#######.##",
				"#.#.#..#..",
				"####.##.##",
				"#.##.##.##",
				"#.#####.##",
				"..#.#.....",
				"##########",
				"#.######.#",
				"#.#####.##",
			},
			expected: 71,
		}, {
			name: "final",
			input: []string{
				"#.#L.L#.##",
				"#LLL#LL.L#",
				"L.#.L..#..",
				"#L##.##.L#",
				"#.#L.LL.LL",
				"#.#L#L#.##",
				"..L.L.....",
				"#L#L##L#L#",
				"#.LLLLLL.L",
				"#.#L#L#.##",
			},
			expected: 37,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			seating := NewSeating()
			for _, i := range tt.input {
				seating.AddRow(i)
			}
			assert.Equal(t, tt.expected, seating.CountOccupied())
		})
	}
}

func TestCheckAdjacent(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		target   rune
		expected int
		i        int
		j        int
	}{
		{
			name: "base",
			input: []string{
				"#.#",
				"###",
				"#.#",
			},
			target:   '#',
			expected: 2,
			i:        1,
			j:        1,
		}, {
			name: "root",
			input: []string{
				"#.#",
				"###",
				"#.#",
			},
			target:   '#',
			expected: 1,
			i:        0,
			j:        0,
		}, {
			name: "corner",
			input: []string{
				"#.#",
				"##L",
				"#L#",
			},
			target:   'L',
			expected: 2,
			i:        2,
			j:        2,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			seating := NewSeating()
			for _, i := range tt.input {
				seating.AddRow(i)
			}
			assert.Equal(t, tt.expected, seating.checkAdjacent(tt.i, tt.j, tt.target))
		})
	}
}

func TestCalculateAdjacencies(t *testing.T) {
	input := []string{
		"L.L",
		"LLL",
		"L.L",
	}
	expected := [][]int{
		{1, 0, 1},
		{3, 2, 3},
		{1, 0, 1},
	}
	seating := NewSeating()
	for _, i := range input {
		seating.AddRow(i)
	}
	seating.CalculateAdjacencies()
	if !reflect.DeepEqual(expected, seating.adjacent) {
		t.Errorf("parsed adjacencies should match")
	}
}

func TestIterate(t *testing.T) {
	input := [][][]rune{
		{
			[]rune("L.LL.LL.LL"),
			[]rune("LLLLLLL.LL"),
			[]rune("L.L.L..L.."),
			[]rune("LLLL.LL.LL"),
			[]rune("L.LL.LL.LL"),
			[]rune("L.LLLLL.LL"),
			[]rune("..L.L....."),
			[]rune("LLLLLLLLLL"),
			[]rune("L.LLLLLL.L"),
			[]rune("L.LLLLL.LL"),
		}, {
			[]rune("#.##.##.##"),
			[]rune("#######.##"),
			[]rune("#.#.#..#.."),
			[]rune("####.##.##"),
			[]rune("#.##.##.##"),
			[]rune("#.#####.##"),
			[]rune("..#.#....."),
			[]rune("##########"),
			[]rune("#.######.#"),
			[]rune("#.#####.##"),
		}, {
			[]rune("#.LL.L#.##"),
			[]rune("#LLLLLL.L#"),
			[]rune("L.L.L..L.."),
			[]rune("#LLL.LL.L#"),
			[]rune("#.LL.LL.LL"),
			[]rune("#.LLLL#.##"),
			[]rune("..L.L....."),
			[]rune("#LLLLLLLL#"),
			[]rune("#.LLLLLL.L"),
			[]rune("#.#LLLL.##"),
		},
	}
	seating := NewSeating()
	seating.seats = input[0]
	seating.CalculateAdjacencies()
	seating.height = len(input[0])
	seating.width = len(input[0][0])
	i := 1
	seating.print()
	for i < len(input) {
		seating.Iterate()
		seating.print()
		if !reflect.DeepEqual(seating.seats, input[i]) {
			t.Errorf("Seating grids should match %v %v", seating.seats, input[i])
		}
		i++
	}
}
