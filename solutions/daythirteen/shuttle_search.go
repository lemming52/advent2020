package daythirteen

import (
	"bufio"
	"fmt"
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

func FindMagicalDepartureTime(shuttles []int) int {
	coefficient := int(1)
	factor := shuttles[0]
	offset := 0
	for i, shuttle := range shuttles[1:] {
		fmt.Println(factor, coefficient, offset, shuttle)
		if shuttle == -1 {
			continue
		}
		factor, coefficient, offset = FindNextShared(
			factor,
			coefficient,
			offset,
			shuttle,
			i+1,
		)
	}
	fmt.Println(factor, coefficient, offset)
	return coefficient*factor - offset

}

func FindNextShared(initialFactor, initialCoeff, initialOffset, nextFactor, nextOffset int) (int, int, int) {
	coeff := initialCoeff
	nextCoefficient := coeffCalculation(initialFactor, coeff, initialOffset, nextFactor, nextOffset)
	for nextCoefficient != float64(int64(nextCoefficient)) {
		coeff++
		nextCoefficient = coeffCalculation(initialFactor, coeff, initialOffset, nextFactor, nextOffset)
	}
	if nextFactor > initialFactor {
		return nextFactor, int(nextCoefficient), nextOffset
	}
	return initialFactor, coeff, initialOffset
}

func coeffCalculation(initialFactor, initialCoeff, initialOffset, nextFactor, nextOffset int) float64 {
	num := initialFactor*initialCoeff - initialOffset + nextOffset
	return float64(num) / float64(nextFactor)
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
	line := scanner.Text()
	shuttles := parseShuttles(line)
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

func parseShuttlesAndSpaces(s string) []int {
	entries := strings.Split(s, ",")
	shuttles := []int{}
	for _, e := range entries {
		if e == "x" {
			shuttles = append(shuttles, -1)
		}
		id, err := strconv.Atoi(e)
		if err != nil {
			log.Fatal(err)
		}
		shuttles = append(shuttles, id)
	}
	return shuttles
}
