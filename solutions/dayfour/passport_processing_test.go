package dayfour

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLine(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Passport
	}{
		{
			name:  "base",
			input: "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
			expected: Passport{
				EyeColour:  "gry",
				PassportID: "860033327",
				ExpiryYear: "2020",
				HairColour: "#fffffd",
				FieldCount: 4,
			},
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			p := Passport{}
			p.ParseLine(tt.input)
			if !reflect.DeepEqual(p, tt.expected) {
				t.Errorf("Passports should match res: %v exp: %v", p, tt.expected)
			}
		})
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		input    Passport
		expected bool
	}{
		{
			name: "A",
			input: Passport{
				EyeColour:  "gry",
				PassportID: "860033327",
				ExpiryYear: "2020",
				HairColour: "#fffffd",
				BirthYear:  "1937",
				IssueYear:  "2017",
				CountryID:  "147",
				Height:     "183cm",
				FieldCount: 8,
			},
			expected: true,
		}, {
			name: "B",
			input: Passport{
				EyeColour:  "amb",
				PassportID: "028048884",
				ExpiryYear: "2023",
				HairColour: "#cfa07d",
				BirthYear:  "1929",
				IssueYear:  "2013",
				CountryID:  "350",
				FieldCount: 7,
			},
			expected: false,
		}, {
			name: "C",
			input: Passport{
				EyeColour:  "brn",
				PassportID: "760753108",
				ExpiryYear: "2024",
				HairColour: "#ae17e1",
				BirthYear:  "1931",
				IssueYear:  "2013",
				Height:     "179cm",
				FieldCount: 7,
			},
			expected: true,
		}, {
			name: "D",
			input: Passport{
				EyeColour:  "brn",
				PassportID: "166559648",
				ExpiryYear: "2025",
				HairColour: "#cfa07d",
				IssueYear:  "2011",
				Height:     "59in",
				FieldCount: 6,
			},
			expected: false,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.input.IsValid(), "validity should match")
		})
	}
}
