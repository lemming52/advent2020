package daythirteen

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// FindEarliest uses the mod of the time w.r.t the bus id to find the wait time, considering all shuttles
func FindEarliest(ids []int, departureTime int) int {
	shuttle := ids[0]
	waitTime := ids[0] - (departureTime % ids[0])
	for _, id := range ids[1:] {
		wait := id - (departureTime % id)
		if wait < waitTime {
			shuttle = id
			waitTime = wait
		}
	}
	return shuttle * waitTime
}

func FindMagicalDepartureTime(shuttles []int) int {
	positions := map[int]int{}
	ids := []int{}
	for i, shuttle := range shuttles {
		positions[shuttle] = i
		ids = append(ids, shuttle)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ids)))
	coeff := 1
	value := ids[0]*coeff - positions[ids[0]]
	valid := TestValidValue(value, &ids, &positions)
	for !valid {
		coeff++
		value = ids[0]*coeff - positions[ids[0]]
		valid = TestValidValue(value, &ids, &positions)
	}
	return value
}

func TestValidValue(value int, ids *[]int, positions *map[int]int) bool {
	for _, id := range (*ids)[1:] {
		if id == -1 {
			return true
		}
		if !isValid(value, id, (*positions)[id]) {
			return false
		}
	}
	return true
}

func isValid(value, id, position int) bool {
	coeff := float64(value+position) / float64(id)
	return coeff == float64(int64(coeff))
}

func coeffCalculation(initialFactor, initialCoeff, initialOffset, nextFactor, nextOffset int) float64 {
	num := initialFactor*initialCoeff - initialOffset + nextOffset
	return float64(num) / float64(nextFactor)
}

func LoadShuttles(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	departureTime, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	scanner.Scan()
	line := scanner.Text()
	shuttles := parseShuttles(line)
	shuttlesAndSpaces := parseShuttlesAndSpaces(line)
	return FindEarliest(shuttles, departureTime), FindMagicalDepartureTime(shuttlesAndSpaces)
}

func parseShuttles(s string) []int {
	entries := strings.Split(s, ",")
	shuttles := []int{}
	for _, e := range entries {
		if e == "x" {
			continue
		}
		id, err := strconv.Atoi(e)
		if err != nil {
			log.Fatal(err)
		}
		shuttles = append(shuttles, id)
	}
	return shuttles
}

func parseShuttlesAndSpaces(s string) []int {
	entries := strings.Split(s, ",")
	shuttles := []int{}
	for _, e := range entries {
		if e == "x" {
			shuttles = append(shuttles, -1)
			continue
		}
		id, err := strconv.Atoi(e)
		if err != nil {
			log.Fatal(err)
		}
		shuttles = append(shuttles, id)
	}
	return shuttles
}
