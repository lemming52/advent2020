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

// ConditionPattern is a regex for the ticket validity conditions, which come in pairs
const ConditionPattern = `([a-z ]+): (\d{2,3})-(\d{2,3}) or (\d{2,3})-(\d{2,3})`

// Validator is a combined struct for validating individual tickets and handling overall stats
type Validator struct {
	values       map[int][]string
	valid        [][]int
	ticket       []int
	fields       map[string]bool
	invalidTotal int
}

// NewValidator instantiates an empty ticket validator
func NewValidator() *Validator {
	return &Validator{
		values:       map[int][]string{},
		valid:        [][]int{},
		ticket:       []int{},
		fields:       map[string]bool{},
		invalidTotal: 0,
	}
}

// AddTicket adds my ticket and initialises the valid values matrix
func (v *Validator) AddTicket(ticket string) {
	values := strings.Split(ticket, ",")
	for _, val := range values {
		entry, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		v.ticket = append(v.ticket, entry)
		v.valid = append(v.valid, []int{})
	}
}

// AddCondition adds a validity check to the validator
func (v *Validator) AddCondition(title, min, max string) {
	v.fields[title] = true
	mn, err := strconv.Atoi(min)
	if err != nil {
		log.Fatal(err)
	}
	mx, err := strconv.Atoi(max)
	if err != nil {
		log.Fatal(err)
	}
	for i := mn; i <= mx; i++ {
		v.values[i] = append(v.values[i], title)
	}
}

// Validate checks if a given ticket is valid, storing it if it is or incrementing the invalid total if not
func (v *Validator) Validate(line string) {
	values := strings.Split(line, ",")
	entries := []int{}
	for _, val := range values {
		entry, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		if !v.isValid(entry) {
			v.invalidTotal += entry
			return
		}
		entries = append(entries, entry)
	}
	for i, e := range entries {
		v.valid[i] = append(v.valid[i], e)
	}
	return
}

func (v *Validator) isValid(val int) bool {
	_, ok := v.values[val]
	return ok
}

func (v *Validator) DetemineProduct() int {
	fields := v.DetermineFields()
	products := map[string]int{}
	product := 1
	for len(fields) > 0 {
		title, index := removeSmallestField(&fields)
		products[title] = index
	}
	fmt.Println(products)
	for k, val := range products {
		if strings.HasPrefix(k, "departure") {
			product *= v.ticket[val]
		}
	}
	return product
}

func (v *Validator) DetermineFields() map[int][]string {
	fieldPositions := map[int][]string{}
	length := len(v.valid[0])
	for i := range v.ticket {
		keys := v.determineField(i, length)
		fieldPositions[i] = keys
	}
	return fieldPositions
}

func (v *Validator) determineField(index, length int) []string {
	fieldCounts := &map[string]int{}
	for _, val := range v.valid[index] {
		fields := v.values[val]
		for _, f := range fields {
			incrementField(fieldCounts, f)
		}
	}
	success := []string{}
	for f, i := range *fieldCounts {
		if i == length {
			success = append(success, f)
		}
	}
	return success
}

func incrementField(fields *map[string]int, f string) {
	_, ok := (*fields)[f]
	if !ok {
		(*fields)[f] = 1
	} else {
		(*fields)[f]++
	}
}

func removeSmallestField(fields *map[int][]string) (string, int) {
	var target int
	var title string
	for i, v := range *fields {
		if len(v) == 1 {
			target = i
			title = v[0]
			delete(*fields, target)
			break
		}
	}
	for i, v := range *fields {
		for j, e := range v {
			if e == title {
				v[len(v)-1], v[j] = v[j], v[len(v)-1]
				(*fields)[i] = v[:len(v)-1]
			}
		}
	}
	return title, target
}

func LoadTickets(path string) (int, int) {
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
			v.AddCondition(match[1], match[2], match[3])
			v.AddCondition(match[1], match[4], match[5])
			continue
		}
		if strings.HasPrefix(line, "your ticket:") {
			scanner.Scan() // My ticket
			v.AddTicket(scanner.Text())
			break
		}
	}
	scanner.Scan() // empty line
	scanner.Scan() // nearby tickets:
	for scanner.Scan() {
		v.Validate(scanner.Text())
	}
	return v.invalidTotal, v.DetemineProduct()
}
