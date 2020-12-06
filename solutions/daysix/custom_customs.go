package daysix

import (
	"bufio"
	"log"
	"os"
)

func TotalQuestions(group []string) (int, int) {
	questions := map[rune]int{}
	for _, person := range group {
		for _, c := range person {
			_, ok := questions[c]
			if !ok {
				questions[c] = 1
			} else {
				questions[c]++
			}
		}
	}
	counter := 0
	people := len(group)
	for _, v := range questions {
		if v == people {
			counter++
		}
	}
	return len(questions), counter
}

func LoadQuestions(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	group := []string{}
	count := 0
	agreementCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			total, agreement := TotalQuestions(group)
			count = count + total
			agreementCount = agreementCount + agreement
			group = []string{}
		} else {
			group = append(group, line)
		}
	}
	total, agreement := TotalQuestions(group)
	count = count + total
	agreementCount = agreementCount + agreement
	return count, agreementCount
}
