package dayeleven

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	// VACANT id of seat
	VACANT = 'L'
	// OCCUPIED id of occupied seat
	OCCUPIED = '#'
	// FLOOR id of floor
	FLOOR = '.'
)

// Seating represents the grid of all seats
type Seating struct {
	seats    [][]rune
	width    int
	height   int
	adjacent [][]int
}

// NewSeating instantiates a new seating grid
func NewSeating() *Seating {
	return &Seating{
		seats:  [][]rune{},
		height: 0,
	}
}

// AddRow adds a row of seats to the grid
func (s *Seating) AddRow(row string) {
	if s.height == 0 {
		s.width = len(row)
	}
	s.seats = append(s.seats, []rune(row))
	s.height++
}

// CountOccupied returns the number of all occupied seats in the grid
func (s *Seating) CountOccupied() int {
	count := 0
	for _, row := range s.seats {
		for _, seat := range row {
			if seat == OCCUPIED {
				count++
			}
		}
	}
	return count
}

// Iterate cycles the state of the seating grid
func (s *Seating) Iterate() bool {
	newGrid := [][]rune{}
	changed := false
	for i, row := range s.seats {
		newRow := []rune{}
		for j, seat := range row {
			switch seat {
			case VACANT:
				if s.checkAdjacent(i, j, OCCUPIED) == 0 {
					newRow = append(newRow, OCCUPIED)
					changed = true
				} else {
					newRow = append(newRow, seat)
				}
			case OCCUPIED:
				if s.checkAdjacent(i, j, OCCUPIED) > 3 {
					newRow = append(newRow, VACANT)
					changed = true
				} else {
					newRow = append(newRow, seat)
				}
			default:
				newRow = append(newRow, seat)
			}
		}
		newGrid = append(newGrid, newRow)
	}
	s.seats = newGrid
	return changed
}

// IterateVisible cycles the state of the seating grid
// using the alternative adjacency rules
func (s *Seating) IterateVisible() bool {
	newGrid := [][]rune{}
	changed := false
	for i, row := range s.seats {
		newRow := []rune{}
		for j, seat := range row {
			switch seat {
			case VACANT:
				if s.checkVisibleAdjacent(i, j, OCCUPIED) == 0 {
					newRow = append(newRow, OCCUPIED)
					changed = true
				} else {
					newRow = append(newRow, seat)
				}
			case OCCUPIED:
				if s.checkVisibleAdjacent(i, j, OCCUPIED) > 4 {
					newRow = append(newRow, VACANT)
					changed = true
				} else {
					newRow = append(newRow, seat)
				}
			default:
				newRow = append(newRow, seat)
			}
		}
		newGrid = append(newGrid, newRow)
	}
	s.seats = newGrid
	return changed
}

// checkAdjacent takes a grid reference and returns the number of adjacent
// seats that match the target rune
func (s *Seating) checkAdjacent(i, j int, target rune) int {
	cells := [][]int{
		{i - 1, j - 1},
		{i - 1, j},
		{i - 1, j + 1},
		{i, j - 1},
		{i, j + 1},
		{i + 1, j - 1},
		{i + 1, j},
		{i + 1, j + 1},
	}
	count := 0
	for _, cell := range cells {
		if cell[0] == -1 || cell[0] == s.height {
			continue
		}
		if cell[1] == -1 || cell[1] == s.width {
			continue
		}
		if s.seats[cell[0]][cell[1]] == target {
			count++
		}
	}
	return count
}

// checkVisibleAdjacent takes a grid reference and returns the number of adjacent
// seats that match the target rune
func (s *Seating) checkVisibleAdjacent(i, j int, target rune) int {
	directions := [][]int{
		{-1, -1}, // NW
		{-1, 0},  // N
		{-1, 1},  // NE
		{0, -1},  // W
		{0, 1},   // E
		{1, -1},  // SW
		{1, 0},   // S
		{1, 1},   // SE
	}
	count := 0
	for _, direction := range directions {
		if s.checkDirection(i, j, direction[0], direction[1], target, FLOOR) {
			count++
		}
	}
	return count
}

// checkDirection takes an increment direction and target and returns
// if the first visible rune matches target, ignoring ignore
func (s *Seating) checkDirection(i, j, xIncrement, yIncrement int, target, ignore rune) bool {
	x := i + xIncrement
	y := j + yIncrement
	for (x != -1 && x != s.height) && (y != -1 && y != s.width) {
		if s.seats[x][y] == target {
			return true
		}
		if s.seats[x][y] != ignore {
			return false
		}
		x += xIncrement
		y += yIncrement
	}
	return false
}

func (s *Seating) print() {
	fmt.Println()
	for _, row := range s.seats {
		fmt.Println(string(row))
	}
}

func LoadSeats(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	seating := NewSeating()
	visibleSeating := NewSeating()
	for scanner.Scan() {
		line := scanner.Text()
		seating.AddRow(line)
		visibleSeating.AddRow(line)
	}
	unstable := true
	for unstable {
		unstable = seating.Iterate()
	}
	unstable = true
	for unstable {
		unstable = visibleSeating.IterateVisible()
	}
	return seating.CountOccupied(), visibleSeating.CountOccupied()
}
