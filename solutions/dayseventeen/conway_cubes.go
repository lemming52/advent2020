package dayseventeen

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

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
	c := NewDimension([]int{turn*2 + 1, len(init[0]) + turn*2, len(init[0]) + turn*2}, 0)
	d := NewDimension([]int{turn*2 + 1, turn*2 + 1, len(init[0]) + turn*2, len(init[0]) + turn*2}, 0)
	c.Initalise(init, []int{turn, turn, turn})
	d.Initalise(init, []int{turn, turn, turn, turn})
	for i := 0; i < turn; i++ {
		c.Iterate()
		d.Iterate()
	}
	return c.Total(), d.Total()
}
