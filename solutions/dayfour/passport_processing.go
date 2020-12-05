package dayfour

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Passport struct {
	BirthYear  string
	IssueYear  string
	ExpiryYear string
	Height     string
	HairColour string
	EyeColour  string
	PassportID string
	CountryID  string
	FieldCount int
}

// ParseLine parses the lines of the input doc into passport fields
func (p *Passport) ParseLine(s string) {
	params := strings.Split(s, " ")
	for _, param := range params {
		components := strings.Split(param, ":")
		key, value := components[0], components[1]
		switch key {
		case "byr":
			p.BirthYear = value
		case "iyr":
			p.IssueYear = value
		case "eyr":
			p.ExpiryYear = value
		case "hgt":
			p.Height = value
		case "hcl":
			p.HairColour = value
		case "ecl":
			p.EyeColour = value
		case "pid":
			p.PassportID = value
		case "cid":
			p.CountryID = value
		}
		p.FieldCount++
	}
}

// IsValid checks if a passport is valid
func (p *Passport) IsValid() bool {
	switch p.FieldCount {
	case 7:
		return p.CountryID == ""
	case 8:
		return true
	default:
		return false
	}
}

// Loadpassports loads all the passports from an input file and checks how many are valid
func LoadPassports(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	passport := Passport{}
	valid := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if passport.IsValid() {
				valid++
			}
			passport = Passport{}
			continue
		}
		passport.ParseLine(line)
	}
	if passport.IsValid() {
		valid++
	}
	return valid
}
