package daythree

import (
	"bufio"
	"log"
	"os"
)

const widthIncrement = 3
const tree = '#'

func TreeCounter(trees []string) int {
	treeWidth := len(trees[0])
	width := 0
	count := 0
	for _, row := range trees {
		if row[width] == tree {
			count++
		}
		width = (width + widthIncrement) % treeWidth
	}
	return count
}

func LoadTrees(path string) int {
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
	return TreeCounter(trees)
}
