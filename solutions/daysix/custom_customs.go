package daysix

import (
	"bufio"
	"log"
	"os"
)

func TotalQuestions(group []string) int {
	questions := map[rune]bool{}
	for _, person := range group {
		for _, c := range person {
			questions[c] = true
		}
	}
	return len(questions)
}

func LoadQuestions(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	group := []string{}
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			count = count + TotalQuestions(group)
			group = []string{}
		} else {
			group = append(group, line)
		}
	}
	count = count + TotalQuestions(group)
	return count
}
