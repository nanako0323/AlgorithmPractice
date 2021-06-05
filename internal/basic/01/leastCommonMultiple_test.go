package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreatestCommonDivisor(t *testing.T) {

	result := greatestCommonDivisor(12, 18)

	assert.Equal(t, 6, result)

	t.Logf("result: %d", result)
}

func TestLeastCommonMultiple(t *testing.T) {

	result := LeastCommonMultiple(12, 18)

	assert.Equal(t, 36, result)

	t.Logf("result: %d", result)
}
