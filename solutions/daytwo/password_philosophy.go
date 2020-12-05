package daytwo

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

const PasswordPattern = `(\d+)-(\d+) ([a-z]){1}: ([a-z]+)`

type Password struct {
	min       int
	max       int
	character rune
	password  string
}

// IsValid applies the rules on the password object to the given password
func (p Password) IsValid() bool {
	if len(p.password) < p.max {
		return false
	}
	count := 0
	for _, c := range p.password {
		if c == p.character {
			count++
			if count > p.max {
				return false
			}
		}
	}
	return count >= p.min
}

// ValidatePasswords loads a file, loads each line into a password object and checks if it is valid
func ValidatePasswords(path string) int {
	/*
		Load each password and ruleset
		check if each password and ruleset is valid
	*/
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pattern, err := regexp.Compile(PasswordPattern)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	valid := 0
	for scanner.Scan() {
		match := pattern.FindStringSubmatch(scanner.Text())
		if match == nil {
			continue
		}
		min, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal(err)
		}
		max, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatal(err)
		}
		password := Password{
			min:       min,
			max:       max,
			character: []rune(match[3])[0],
			password:  match[4],
		}
		if password.IsValid() {
			valid++
		}
	}
	return valid
}
