package dayeleven

import "fmt"

const (
	VACANT   = 'L'
	OCCUPIED = '#'
	FLOOR    = '.'
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
		seats:    [][]rune{},
		height:   0,
		adjacent: [][]int{},
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

// CalculateAdjacencies is a convencience function that instantiates a
// number of adjacent seats for all locations
func (s *Seating) CalculateAdjacencies() {
	for i, row := range s.seats {
		vals := []int{}
		for j := range row {
			if s.seats[i][j] != FLOOR {
				vals = append(vals, s.checkAdjacent(i, j, VACANT))
			} else {
				vals = append(vals, 0)
			}
		}
		s.adjacent = append(s.adjacent, vals)
	}

}

// checkAdjacent takes a grid reference and returns the number of adjacent
// seats that match the target rune
func (s *Seating) checkAdjacent(i, j int, target rune) int {
	adjacent := []rune{}
	if i > 0 {
		adjacent = append(adjacent, s.seats[i-1][j])
	}
	if i < s.height-1 {
		adjacent = append(adjacent, s.seats[i+1][j])
	}
	if j > 0 {
		adjacent = append(adjacent, s.seats[i][j-1])
	}
	if j < s.width-1 {
		adjacent = append(adjacent, s.seats[i][j+1])
	}
	count := 0
	for _, r := range adjacent {
		if r == target {
			count++
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
				if s.adjacent[i][j] == 0 || s.checkAdjacent(i, j, OCCUPIED) == 0 {
					newRow = append(newRow, OCCUPIED)
					changed = true
				} else {
					newRow = append(newRow, seat)
				}
			case OCCUPIED:
				if s.adjacent[i][j] > 3 && s.checkAdjacent(i, j, OCCUPIED) > 3 {
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

func (s *Seating) print() {
	fmt.Println()
	for _, row := range s.seats {
		fmt.Println(string(row))
	}
}

func LoadSeats(path string) int {
	return 0
}
