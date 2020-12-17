package daythirteen

import (
	"bufio"
	"log"
	"os"
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

func LoadShuttles(path string) int {
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
	shuttles := parseShuttles(scanner.Text())
	return FindEarliest(shuttles, departureTime)
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
