package dayfifteen

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	numbers map[int]int
	turn    int
}

func (g *Game) Initialise(starting []int) int {
	g.numbers = map[int]int{}
	var val int
	for _, n := range starting {
		val = g.Iterate(n)
	}
	return val
}

func (g *Game) Iterate(spoken int) int {
	g.turn++
	count, ok := g.numbers[spoken]
	if !ok {
		g.numbers[spoken] = g.turn
		return 0
	}
	next := g.turn - count
	g.numbers[spoken] = g.turn
	return next
}

func Play(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	entries := strings.Split(line, ",")
	starting := []int{}
	for _, i := range entries {
		v, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		starting = append(starting, v)
	}
	g := Game{}
	val := g.Initialise(starting)
	i := 0
	for i < 2020-len(starting)-1 {
		val = g.Iterate(val)
		i++
	}
	first := val
	for i < 30000000-len(starting)-1 {
		val = g.Iterate(val)
		i++
	}
	return first, val
}
