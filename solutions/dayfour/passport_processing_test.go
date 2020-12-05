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
				ExpiryYear: 2020,
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
				ExpiryYear: 2020,
				HairColour: "#fffffd",
				BirthYear:  1937,
				IssueYear:  2017,
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
				ExpiryYear: 2023,
				HairColour: "#cfa07d",
				BirthYear:  1929,
				IssueYear:  2013,
				CountryID:  "350",
				FieldCount: 7,
			},
			expected: false,
		}, {
			name: "C",
			input: Passport{
				EyeColour:  "brn",
				PassportID: "760753108",
				ExpiryYear: 2024,
				HairColour: "#ae17e1",
				BirthYear:  1931,
				IssueYear:  2013,
				Height:     "179cm",
				FieldCount: 7,
			},
			expected: true,
		}, {
			name: "D",
			input: Passport{
				EyeColour:  "brn",
				PassportID: "166559648",
				ExpiryYear: 2025,
				HairColour: "#cfa07d",
				IssueYear:  2011,
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

func TestIsExtraValid(t *testing.T) {
	tests := []struct {
		name     string
		input    Passport
		expected bool
	}{
		{
			name: "A",
			input: Passport{
				EyeColour:  "grn",
				PassportID: "087499704",
				ExpiryYear: 2030,
				HairColour: "#623a2f",
				BirthYear:  1980,
				IssueYear:  2012,
				Height:     "74in",
				FieldCount: 7,
			},
			expected: true,
		}, {
			name: "B",
			input: Passport{
				EyeColour:  "blu",
				PassportID: "896056539",
				ExpiryYear: 2029,
				HairColour: "#a97842",
				BirthYear:  1989,
				IssueYear:  2014,
				CountryID:  "129",
				Height:     "165cm",
				FieldCount: 8,
			},
			expected: true,
		}, {
			name: "C",
			input: Passport{
				EyeColour:  "hzl",
				PassportID: "545766238",
				ExpiryYear: 2022,
				HairColour: "#888785",
				BirthYear:  2001,
				IssueYear:  2015,
				CountryID:  "88",
				Height:     "164cm",
				FieldCount: 8,
			},
			expected: true,
		}, {
			name: "D",
			input: Passport{
				EyeColour:  "blu",
				PassportID: "093154719",
				ExpiryYear: 2021,
				HairColour: "#b6652a",
				BirthYear:  1944,
				IssueYear:  2010,
				Height:     "158cm",
				FieldCount: 7,
			},
			expected: true,
		}, {
			name: "E",
			input: Passport{
				EyeColour:  "amb",
				PassportID: "186cm",
				ExpiryYear: 1972,
				HairColour: "#18171d",
				BirthYear:  1926,
				IssueYear:  2018,
				Height:     "170",
				CountryID:  "100",
				FieldCount: 8,
			},
			expected: false,
		}, {
			name: "F",
			input: Passport{
				EyeColour:  "grn",
				PassportID: "012533040",
				ExpiryYear: 1967,
				HairColour: "#602927",
				BirthYear:  1946,
				IssueYear:  2019,
				Height:     "170cm",
				FieldCount: 7,
			},
			expected: false,
		}, {
			name: "G",
			input: Passport{
				EyeColour:  "brn",
				PassportID: "021572410",
				ExpiryYear: 2020,
				HairColour: "dab227",
				BirthYear:  1992,
				IssueYear:  2012,
				Height:     "182cm",
				CountryID:  "227",
				FieldCount: 8,
			},
			expected: false,
		}, {
			name: "H",
			input: Passport{
				EyeColour:  "zzz",
				PassportID: "3556412378",
				ExpiryYear: 2038,
				HairColour: "74454a",
				BirthYear:  2007,
				IssueYear:  2023,
				Height:     "59cm",
				CountryID:  "227",
				FieldCount: 8,
			},
			expected: false,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.input.IsExtraValid(), "validity should match")
		})
	}
}
