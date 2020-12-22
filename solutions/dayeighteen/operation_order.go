package dayeighteen

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"unicode"
)

const (
	ADDITION       = iota
	SUBTRACTION    = iota
	MULTIPLICATION = iota
	DIVISION       = iota
)

func Parse(s string) int {
	components := []int{}
	i := 0
	for i < len(s) {
		character := rune(s[i])
		if character == ' ' {
			i++
			continue
		}
		if unicode.IsDigit(character) {
			j := i + 1
			for j < len(s) && unicode.IsDigit(rune(s[j])) {
				j++
			}
			var substring string
			if j == len(s) {
				substring = s[i:]
			} else {
				substring = s[i:j]
			}
			val, err := strconv.Atoi(substring)
			if err != nil {
				log.Fatal(err)
			}
			components = append(components, val)
			i = j
		}
		switch character {
		case '+':
			components = append(components, ADDITION)
			i++
		case '-':
			components = append(components, SUBTRACTION)
			i++
		case '*':
			components = append(components, MULTIPLICATION)
			i++
		case '/':
			components = append(components, DIVISION)
			i++
		case '(':
			j := i + 1
			openCount := 1
			for openCount > 0 {
				switch rune(s[j]) {
				case ')':
					openCount--
				case '(':
					openCount++
				}
				j++
			}

			val := Parse(s[i+1 : j-1])
			components = append(components, val)
			i = j + 1
		}
	}
	return Compute(components)
}

func Compute(components []int) int {
	odd := true
	total := components[0]
	operation := -1
	for _, c := range components[1:] {
		if odd {
			operation = c
			odd = false
		} else {
			switch operation {
			case ADDITION:
				total += c
			case SUBTRACTION:
				total -= c
			case MULTIPLICATION:
				total *= c
			case DIVISION:
				total /= c
			}
			odd = true
		}
	}
	return total
}

func LoadOperations(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		total += Parse(scanner.Text())
	}
	return total
}
