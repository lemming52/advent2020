package daynine

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func FindOutlier(numbers []int, window int) int {
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

// FindOutlierSum iterates across the list to find a continuous set of
// numbers that sum to the target, and returns the smallest and largest number
func FindOutlierSum(numbers []int, target int) int {
	i, j := 0, 1
	sum := numbers[i]
	for j < len(numbers) {
		if sum == target {
			subset := numbers[i:j]
			sort.Ints(subset)
			return subset[len(subset)-1] + subset[0]
		} else if (sum < target) || (i == j-1) {
			sum += numbers[j]
			j++
		} else {
			sum -= numbers[i]
			i++
		}
	}
	return 0
}

func LoadNumbers(path string) (int, int) {
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
	outlier := FindOutlier(numbers, 25)
	return outlier, FindOutlierSum(numbers, outlier)
}
