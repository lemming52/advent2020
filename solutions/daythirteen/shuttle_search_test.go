package daythirteen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindEarliest(t *testing.T) {
	ids := []int{7, 13, 59, 31, 19}
	departureTime := 939
	expected := 295
	res := FindEarliest(ids, departureTime)
	assert.Equal(t, expected, res)
}
