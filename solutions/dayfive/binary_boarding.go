package dayfive

import (
	"bufio"
	"log"
	"os"
)

const (
	forward = 'F'
	left    = 'L'
)

func ParseBoardingPass(s string) (int, int) {
	rows := s[:7]
	columns := s[7:]
	row := CalculateValue(rows, 128, forward)
	column := CalculateValue(columns, 8, left)
	return row, column
}

func CalculateValue(s string, max int, marker rune) int {
	val := max
	delta := max / 2
	for _, c := range s {
		if c == marker {
			val = val - delta
		}
		delta = delta / 2
	}
	return val - 1
}

func LoadBoardingPasses(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	max := 0
	for scanner.Scan() {
		row, column := ParseBoardingPass(scanner.Text())
		value := row*8 + column
		if value > max {
			max = value
		}
	}
	return max
}
