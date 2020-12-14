package dayten

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// CountAdapterIntervals takes a sorted list of adapters, prepended by a 0 value
// and computes the product of the number of 1 and 3V jumps, including the extra
// final jump to the device
func CountAdapterIntervals(adapters []int) int {
	counts := map[int]int{
		1: 0, 2: 0, 3: 1, // Know last jump is 3V
	}
	for i, a := range adapters[:len(adapters)-1] {
		delta := adapters[i+1] - a
		counts[delta]++
	}
	return counts[1] * counts[3]
}

// DetermineValid Values
func DetermineValidCombinations(adapters []int) int {
	/*
		Figure out smallest path, return 2*(len(numbers not in path))
	*/
	subsets := [][]int{}
	i := 0
	total := 1
	penultimate := len(adapters) - 1
	for i < penultimate {
		nextOptional := findOptional(adapters[i:penultimate]) + i
		nextRequired := findRequired(adapters[nextOptional:penultimate]) + nextOptional
		if nextOptional != penultimate {
			subset := adapters[nextOptional-1 : nextRequired+1]
			subsets = append(subsets, subset)
		}
		i = nextRequired
	}
	for _, subset := range subsets {
		fmt.Println(subset)
		total *= calculateCombinations(subset)
	}
	return total
}

func findOptional(adapters []int) int {
	i := 1
	for i < len(adapters)-1 {
		if adapters[i+1]-adapters[i-1] >= 4 {
			i++
			continue
		}
		return i
	}
	return len(adapters)
}

func findRequired(adapters []int) int {
	i := 1
	for i < len(adapters)-1 {
		if adapters[i+1]-adapters[i-1] >= 4 {
			return i
		}
		i++
	}
	return len(adapters)
}

func calculateCombinations(subset []int) int {
	values := map[int]int{
		3: 2,
		4: 4,
		5: 7,
		6: 11,
	}
	return values[len(subset)]
}

func LoadAdapters(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	adapters := []int{0} // Initial value
	for scanner.Scan() {
		// Sort on read for efficiency.
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		i := sort.SearchInts(adapters, val)
		adapters = append(adapters, 0)
		copy(adapters[i+1:], adapters[i:])
		adapters[i] = val
	}
	return CountAdapterIntervals(adapters), DetermineValidCombinations(adapters)
}
