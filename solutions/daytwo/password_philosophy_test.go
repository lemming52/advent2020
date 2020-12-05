package daytwo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordIsValid(t *testing.T) {
	tests := []struct {
		name     string
		password Password
		expected bool
	}{
		{
			name: "base",
			password: Password{
				min:       1,
				max:       3,
				character: 'a',
				password:  "abcde",
			},
			expected: true,
		}, {
			name: "base2",
			password: Password{
				min:       1,
				max:       3,
				character: 'b',
				password:  "cdefg",
			},
			expected: false,
		}, {
			name: "base3",
			password: Password{
				min:       2,
				max:       9,
				character: 'c',
				password:  "ccccccccc",
			},
			expected: true,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res := tt.password.IsValid()
			assert.Equal(t, tt.expected, res, "password validity wrong")
		})
	}
}

func TestPasswordIsAlternativeValid(t *testing.T) {
	tests := []struct {
		name     string
		password Password
		expected bool
	}{
		{
			name: "base",
			password: Password{
				min:       1,
				max:       3,
				character: 'a',
				password:  "abcde",
			},
			expected: true,
		}, {
			name: "base2",
			password: Password{
				min:       1,
				max:       3,
				character: 'b',
				password:  "cdefg",
			},
			expected: false,
		}, {
			name: "base3",
			password: Password{
				min:       2,
				max:       9,
				character: 'c',
				password:  "ccccccccc",
			},
			expected: false,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res := tt.password.IsValidDownTheRoad()
			assert.Equal(t, tt.expected, res, "password validity wrong")
		})
	}
}
