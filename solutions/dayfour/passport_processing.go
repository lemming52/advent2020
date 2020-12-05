package dayfour

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

const HairColourPattern = `#[0-9a-f]{6}`

type Passport struct {
	BirthYear  int
	IssueYear  int
	ExpiryYear int
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
			value, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			p.BirthYear = value
		case "iyr":
			value, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			p.IssueYear = value
		case "eyr":
			value, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
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

// IsExtraValid has more stringent checks on the passport
func (p *Passport) IsExtraValid() bool {
	if p.FieldCount < 7 {
		return false
	}
	if p.BirthYear < 1920 || p.BirthYear > 2002 {
		return false
	}
	if p.IssueYear < 2010 || p.IssueYear > 2020 {
		return false
	}
	if p.ExpiryYear < 2020 || p.ExpiryYear > 2030 {
		return false
	}
	if !p.IsHeightValid() {
		return false
	}
	if !p.IsHairColourValid() {
		return false
	}
	if !p.IsEyeColourValid() {
		return false
	}
	return p.IsPassportIDValid()
}

// IsHeightValid Checks if the height is valid
func (p *Passport) IsHeightValid() bool {
	length := len(p.Height)
	if length < 3 {
		return false
	}
	unit := p.Height[length-2:]
	value, err := strconv.Atoi(p.Height[:length-2])
	if err != nil {
		log.Fatal(err)
	}
	switch unit {
	case "cm":
		return value <= 193 && value >= 150
	case "in":
		return value <= 76 && value >= 59
	}
	return false
}

// IsHairColourValid checks the hair colour of the passport
func (p *Passport) IsHairColourValid() bool {
	pattern, err := regexp.Compile(HairColourPattern)
	if err != nil {
		log.Fatal(err)
	}
	return pattern.MatchString(p.HairColour)
}

// IsEyeColourValid checks the eye colour of the passport
func (p *Passport) IsEyeColourValid() bool {
	colours := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
	_, ok := colours[p.EyeColour]
	return ok
}

// IsPassportIDValid checks the passportID of the passport
func (p *Passport) IsPassportIDValid() bool {
	for i, c := range p.PassportID {
		if !unicode.IsDigit(c) {
			return false
		}
		if i >= 9 {
			return false
		}
	}
	return len(p.PassportID) == 9
}

// LoadPassports loads all the passports from an input file and checks how many are valid
func LoadPassports(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	passport := Passport{}
	valid, extraValid := 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if passport.IsValid() {
				valid++
			}
			if passport.IsExtraValid() {
				extraValid++
			}
			passport = Passport{}
			continue
		}
		passport.ParseLine(line)
	}
	if passport.IsValid() {
		valid++
	}
	if passport.IsExtraValid() {
		extraValid++
	}
	return valid, extraValid
}
