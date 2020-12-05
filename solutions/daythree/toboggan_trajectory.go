package daythree

import (
	"bufio"
	"log"
	"os"
)

const tree = '#'

// TreeCounter counts the number of trees encountered traversing the trees at a given slope
func TreeCounter(trees []string, widthIncrement, heightIncrement int) int {
	treeWidth := len(trees[0])
	totalHeight := len(trees)
	width, height, count := 0, 0, 0

	for height < totalHeight {
		row := trees[height]
		if row[width] == tree {
			count++
		}
		width = (width + widthIncrement) % treeWidth
		height = height + heightIncrement
	}
	return count
}

//  LoadTrees recovers the tree definition and then performs count calculations for certain slopes
func LoadTrees(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	trees := []string{}
	for scanner.Scan() {
		trees = append(trees, scanner.Text())
	}
	a := TreeCounter(trees, 3, 1)

	slopes := [][]int{
		{1, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	b := a
	for _, slope := range slopes {
		b = b * TreeCounter(trees, slope[0], slope[1])
	}
	return a, b
}
