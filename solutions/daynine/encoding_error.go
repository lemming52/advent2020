package daynine

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ParseNumbers(numbers []int, window int) int {
	i, j := 0, window
	for j < len(numbers) {
		if !TwoSum(numbers[i:j], numbers[j]) {
			return numbers[j]
		}
		i++
		j++
	}
	return 0
}

// TwoSum is an implementation of the earlier problem to check for two sum
func TwoSum(numbers []int, target int) bool {
	complements := map[int]bool{}
	for _, n := range numbers {
		complement := target - n
		_, ok := complements[complement]
		if ok {
			return true
		}
		complements[n] = true
	}
	return false
}

func LoadNumbers(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbers := []int{}
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, val)
	}
	return ParseNumbers(numbers, 25)
}
