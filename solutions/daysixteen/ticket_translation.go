package daysixteen

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const ConditionPattern = `[a-z ]+: (\d{2,3})-(\d{2,3}) or (\d{2,3})-(\d{2,3})`

type Validator struct {
	values       map[int]bool
	invalidTotal int
}

func NewValidator() *Validator {
	return &Validator{
		values:       map[int]bool{},
		invalidTotal: 0,
	}
}

func (v *Validator) AddCondition(min, max string) {
	mn, err := strconv.Atoi(min)
	if err != nil {
		log.Fatal(err)
	}
	mx, err := strconv.Atoi(max)
	if err != nil {
		log.Fatal(err)
	}
	for i := mn; i <= mx; i++ {
		v.values[i] = true
	}
}

func (v *Validator) Validate(line string) {
	values := strings.Split(line, ",")
	for _, val := range values {
		entry, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		if !v.isValid(entry) {
			v.invalidTotal += entry
			return
		}
	}
	return
}

func (v *Validator) isValid(val int) bool {
	_, ok := v.values[val]
	return ok
}

func LoadTickets(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	condition, err := regexp.Compile(ConditionPattern)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	v := NewValidator()
	// Scan conditions
	for scanner.Scan() {
		line := scanner.Text()
		match := condition.FindStringSubmatch(line)
		if match != nil {
			v.AddCondition(match[1], match[2])
			v.AddCondition(match[3], match[4])
			continue
		}
		if strings.HasPrefix(line, "your ticket:") {
			scanner.Scan() // My ticket
			break
		}
	}
	scanner.Scan() // empty line
	scanner.Scan() // nearby tickets:
	fmt.Println(scanner.Text())
	for scanner.Scan() {
		v.Validate(scanner.Text())
	}
	return v.invalidTotal
}
