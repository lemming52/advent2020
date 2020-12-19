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

// FindMagicalDepartureTime is the function for solving the niche case of the part 2
func FindMagicalDepartureTime(shuttles []int) int {
	positions := map[int]int{}
	ids := []int{}
	for i, shuttle := range shuttles {
		positions[shuttle] = i
		ids = append(ids, shuttle)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ids)))
	N := 1
	for _, id := range ids {
		if id == -1 {
			break
		}
		N *= id
	}

	total := 0
	for _, id := range ids {
		if id == -1 {
			break
		}
		total += ChineseRemainderComponent(id, positions[id], N)
	}
	return total % N
}

// ChineseRemainderComponent computes the term for remained
func ChineseRemainderComponent(id, position, N int) int {
	/*
		derived from https://www.dave4math.com/mathematics/chinese-remainder-theorem/
	*/
	n := N / id
	u := 1
	val := float64(u*n+1) / float64(id)
	for val != float64(int64(val)) {
		u++
		val = float64(u*n+1) / float64(id)
	}
	return u * n * position
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
