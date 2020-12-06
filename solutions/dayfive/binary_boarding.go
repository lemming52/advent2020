package dayfive

import (
	"bufio"
	"log"
	"os"
)

const (
	forward  = 'F'
	left     = 'L'
	midpoint = 512
)

type SeatNode struct {
	rowMin    int
	rowMax    int
	colMin    int
	colMax    int
	nodeDepth int
	rowDelta  int
	colDelta  int
	Children  int
	id        int
	lowerNode *SeatNode
	upperNode *SeatNode
}

func (n *SeatNode) AddSeat(s string) {
	if s == "" {
		n.id = n.rowMin*8 + n.colMin
		return
	}
	n.Children++
	if n.nodeDepth < 7 {
		if s[0] == forward {
			if n.lowerNode == nil {
				n.lowerNode = &SeatNode{
					rowMin:    n.rowMin,
					rowMax:    n.rowMax - n.rowDelta,
					colMin:    0,
					colMax:    8,
					nodeDepth: n.nodeDepth + 1,
					rowDelta:  n.rowDelta / 2,
					colDelta:  4,
					Children:  0,
					id:        -1,
				}
			}
			n.lowerNode.AddSeat(s[1:])
		} else {
			if n.upperNode == nil {
				n.upperNode = &SeatNode{
					rowMin:    n.rowMin + n.rowDelta,
					rowMax:    n.rowMax,
					colMin:    0,
					colMax:    8,
					nodeDepth: n.nodeDepth + 1,
					rowDelta:  n.rowDelta / 2,
					colDelta:  4,
					Children:  0,
					id:        -1,
				}
			}
			n.upperNode.AddSeat(s[1:])
		}
	} else {
		if s[0] == left {
			if n.lowerNode == nil {
				n.lowerNode = &SeatNode{
					rowMin:    n.rowMin,
					rowMax:    n.rowMax,
					colMin:    n.colMin,
					colMax:    n.colMax - n.colDelta,
					nodeDepth: n.nodeDepth + 1,
					rowDelta:  0,
					colDelta:  n.colDelta / 2,
					Children:  0,
					id:        -1,
				}
			}
			n.lowerNode.AddSeat(s[1:])
		} else {
			if n.upperNode == nil {
				n.upperNode = &SeatNode{
					rowMin:    n.rowMin,
					rowMax:    n.rowMax,
					colMin:    n.colMin + n.colDelta,
					colMax:    n.colMax,
					nodeDepth: n.nodeDepth + 1,
					rowDelta:  0,
					colDelta:  n.colDelta / 2,
					Children:  0,
					id:        -1,
				}
			}
			n.upperNode.AddSeat(s[1:])
		}
	}
	return
}

func (n *SeatNode) FindMin() int {
	if n.id != -1 {
		return n.id
	}
	if n.lowerNode != nil {
		return n.lowerNode.FindMin()
	}
	return n.upperNode.FindMin()
}

func (n *SeatNode) FindMax() int {
	if n.id != -1 {
		return n.id
	}
	if n.upperNode != nil {
		return n.upperNode.FindMax()
	}
	return n.lowerNode.FindMax()
}

func (n *SeatNode) FindSpace(midpoint, delta, min, max int) int {
	if n.lowerNode == nil {
		return n.upperNode.id - 1
	}
	if n.upperNode == nil {
		return n.lowerNode.id + 1
	}
	expectedLower := midpoint - min
	expectedUpper := max - midpoint + 1
	if n.lowerNode.Children == expectedLower {
		return n.upperNode.FindSpace(midpoint+delta, delta/2, midpoint, max)
	}
	if n.upperNode.Children == expectedUpper {
		return n.lowerNode.FindSpace(midpoint-delta, delta/2, min, midpoint-1)
	}
	return 0
}

func LoadBoardingPasses(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	root := SeatNode{
		rowMin:    0,
		rowMax:    127,
		colMin:    0,
		colMax:    8,
		nodeDepth: 0,
		rowDelta:  64,
		colDelta:  4,
		Children:  0,
		id:        -1,
	}

	scanner := bufio.NewScanner(file)
	counter := 0
	for scanner.Scan() {
		counter++
		seat := scanner.Text()
		root.AddSeat(seat)
	}
	max := root.FindMax()
	min := root.FindMin()
	return max, root.FindSpace(512, 256, min, max)
}
