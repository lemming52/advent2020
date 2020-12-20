package dayseventeen

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func LoadCube(path string) int {
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
	offsetX, offsetY := turn, turn
	c.Initalise(init, offsetX, offsetY)
	for i := 0; i < turn; i++ {
		c.Iterate()
	}
	return c.Total()
}
