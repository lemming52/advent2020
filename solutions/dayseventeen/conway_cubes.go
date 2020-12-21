package dayseventeen

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Cube struct {
	dimensions int
	grid       [][][]bool
	turn       int
	X          int
	Y          int
	Z          int
}

func NewCube(order, z, x, y int) *Cube {
	grid := [][][]bool{}
	for i := 0; i < z; i++ {
		plane := [][]bool{}
		for j := 0; j < x; j++ {
			row := []bool{}
			for k := 0; k < y; k++ {
				row = append(row, false)
			}
			plane = append(plane, row)
		}
		grid = append(grid, plane)
	}
	return &Cube{
		grid:       grid,
		dimensions: order,
		turn:       0,
		X:          x,
		Y:          y,
		Z:          z,
	}
}

func (c *Cube) Initalise(slice [][]bool, x, y int) {
	for i, row := range slice {
		for j, entry := range row {
			c.grid[c.Z/2][i+x][j+y] = entry
		}
	}
}

func (c *Cube) Iterate() {
	adjustments := [][]int{}
	for i, plane := range c.grid {
		for j, row := range plane {
			for k, val := range row {
				count := c.checkAdjacent(i, j, k)
				if val { // active
					if count < 2 || count > 3 {
						adjustments = append(adjustments, []int{i, j, k, -1})
					}
				} else { // inactive
					if count == 3 {
						adjustments = append(adjustments, []int{i, j, k, 1})
					}
				}
			}
		}
	}

	for _, a := range adjustments {
		c.grid[a[0]][a[1]][a[2]] = a[3] == 1
	}
}

func (c *Cube) checkAdjacent(i, j, k int) int {
	cells := [][]int{
		{i - 1, j - 1, k - 1}, {i - 1, j, k - 1}, {i - 1, j + 1, k - 1},
		{i, j - 1, k - 1}, {i, j, k - 1}, {i, j + 1, k - 1},
		{i + 1, j - 1, k - 1}, {i + 1, j, k - 1}, {i + 1, j + 1, k - 1},
		{i - 1, j - 1, k}, {i - 1, j, k}, {i - 1, j + 1, k},
		{i, j - 1, k}, {i, j + 1, k},
		{i + 1, j - 1, k}, {i + 1, j, k}, {i + 1, j + 1, k},
		{i - 1, j - 1, k + 1}, {i - 1, j, k + 1}, {i - 1, j + 1, k + 1},
		{i, j - 1, k + 1}, {i, j, k + 1}, {i, j + 1, k + 1},
		{i + 1, j - 1, k + 1}, {i + 1, j, k + 1}, {i + 1, j + 1, k + 1},
	}
	count := 0
	for _, cell := range cells {
		if cell[0] == -1 || cell[0] == c.Z {
			continue
		}
		if cell[1] == -1 || cell[1] == c.X {
			continue
		}
		if cell[2] == -1 || cell[2] == c.Y {
			continue
		}
		if c.grid[cell[0]][cell[1]][cell[2]] {
			count++
		}
	}
	return count
}

func (c *Cube) print() {
	for i, plane := range c.grid {
		fmt.Println(fmt.Sprintf("plane: %d", i))
		for _, row := range plane {
			r := []string{}
			for _, entry := range row {
				if entry {
					r = append(r, "#")
				} else {
					r = append(r, ".")
				}
			}
			fmt.Println(r)
		}
	}
}

func (c *Cube) Total() int {
	count := 0
	for _, plane := range c.grid {
		for _, row := range plane {
			for _, val := range row {
				if val {
					count++
				}
			}
		}
	}
	return count
}

type Dimension struct {
	entries  *[]*Dimension
	order    int
	position int
	value    bool
	space    [][]int
}

func NewDimension(dimensions []int, position int) *Dimension {
	if dimensions == nil {
		return &Dimension{
			order:    0,
			position: position,
			value:    false,
		}
	}
	entries := []*Dimension{}
	for i := 0; i < dimensions[0]; i++ {
		if len(dimensions) == 1 {
			entries = append(entries, NewDimension(nil, i))
		} else {
			entries = append(entries, NewDimension(dimensions[1:], i))
		}
	}
	return &Dimension{
		order:    len(dimensions),
		position: position,
		entries:  &entries,
	}
}

func (d *Dimension) Set(coords []int, val bool) {
	if len(coords) == 1 {
		(*d.entries)[coords[0]].value = val
		return
	}
	(*d.entries)[coords[0]].Set(coords[1:], val)
}

func (d *Dimension) Initalise(init [][]bool, offset []int) {
	for i, row := range init {
		for j, val := range row {
			length := len(offset)
			newOffset := make([]int, length)
			copy(newOffset, offset)
			newOffset[length-2] += i
			newOffset[length-1] += j
			d.Set(newOffset, val)
		}
	}
	d.space = d.GetCoords()
}

func (d *Dimension) CheckAdjacent(coords []int) int {
	total := 0
	if len(coords) == 1 {
		for i := -1; i <= 1; i++ {
			index := coords[0] + i
			if index > 0 && index < len(*d.entries) && (*d.entries)[index].value {
				total++
			}
		}
		return total
	}
	for i := -1; i <= 1; i++ {
		index := coords[0] + i
		if index > 0 && index < len(*d.entries) {
			total += (*d.entries)[index].CheckAdjacent(coords[1:])
		}
	}
	return total
}

func (d *Dimension) Get(coords []int) bool {
	if len(coords) == 1 {
		return (*d.entries)[coords[0]].value
	}
	return (*d.entries)[coords[0]].Get(coords[1:])
}

func (d *Dimension) GetCoords() [][]int {
	if d.order == 1 {
		elems := [][]int{}
		for i := range *d.entries {
			elems = append(elems, []int{i})
		}
		return elems
	}
	elems := [][]int{}
	for i, elem := range *d.entries {
		coords := elem.GetCoords()
		for _, e := range coords {
			coord := append([]int{i}, e...)
			elems = append(elems, coord)
		}
	}
	return elems
}

func (d *Dimension) Iterate() {
	adjustments := [][]int{}
	for _, coords := range d.space {
		value := d.Get(coords)
		count := d.CheckAdjacent(coords)
		if value { // active
			count--
			if count < 2 || count > 3 {
				adjustments = append(adjustments, append(coords, -1))
			}
		} else { // inactive
			if count == 3 {
				adjustments = append(adjustments, append(coords, 1))
			}
		}
	}
	length := len(d.space[0])
	for _, a := range adjustments {
		d.Set(a[:length], a[length] == 1)
	}
}

func (d *Dimension) Total() int {
	total := 0
	for _, coords := range d.space {
		if d.Get(coords) {
			total++
		}
	}
	return total
}

func (d *Dimension) print() string {
	switch d.order {
	case 0:
		if d.value {
			return "#"
		}
		return "."
	case 1:
		cells := []string{}
		for _, elem := range *d.entries {
			cells = append(cells, elem.print())
		}
		return strings.Join(cells[:], "")
	case 2:
		for _, elem := range *d.entries {
			fmt.Println(elem.print())
		}
		return ""
	default:
		for i, elem := range *d.entries {
			fmt.Println(fmt.Sprintf("Order: %d = %d", d.order, i))
			elem.print()
		}
		return ""
	}
}

func LoadCube(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	init := [][]bool{}
	for scanner.Scan() {
		row := []bool{}
		for _, c := range scanner.Text() {
			row = append(row, c == '#')
		}
		init = append(init, row)
	}
	turn := 6
	c := NewCube(3, turn*2+1, len(init[0])+turn*2, len(init[0])+turn*2)
	d := NewDimension([]int{turn*2 + 1, turn*2 + 1, len(init[0]) + turn*2, len(init[0]) + turn*2}, 0)
	offsetX, offsetY := turn, turn
	c.Initalise(init, offsetX, offsetY)
	d.Initalise(init, []int{turn, turn, turn, turn})
	for i := 0; i < turn; i++ {
		c.Iterate()
		d.Iterate()
	}
	return c.Total(), d.Total()
}
