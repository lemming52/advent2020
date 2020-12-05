package dayone

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

// SUM is the desired number for two entries to sum to
const SUM = 2020

// RepairReport Takes a list of numbers, finds the two entries that sum up to 2020
// and returns the product of the two numbers
// So, two sum
func RepairReport(numbers []int64) int64 {

	complements := map[int64]bool{}
	for _, n := range numbers {
		complement := SUM - n
		_, ok := complements[complement]
		if ok {
			return n * complement
		}
		complements[n] = true
	}
	return 0
}

// RepairReportExtra finds three numbers that add to the sum in the list
// returning that
// so, three sum
func RepairReportExtra(numbers []int64) int64 {
	sort.Slice(numbers, func(i, j int) bool { return numbers[i] < numbers[j] })
	length := len(numbers)
	for i, n := range numbers {
		if n > SUM {
			return 0
		}
		j := i + 1
		k := length - 1
		for j < k {
			b := numbers[j]
			c := numbers[k]
			total := n + b + c
			if total == SUM {
				return n * b * c
			}
			if total > SUM {
				k--
			} else {
				j++
			}
		}
	}
	return 0
}

// LoadReport loads an input text file and executes the repair report function
func LoadReport(reportPath string) (int64, int64) {
	file, err := os.Open(reportPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	reportNumbers := []int64{}
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		reportNumbers = append(reportNumbers, int64(n))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return RepairReport(reportNumbers), RepairReportExtra(reportNumbers)
}
