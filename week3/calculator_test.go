package week3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalc(t *testing.T) {
	assert.Equal(t, 42, calculate("42"))
	assert.Equal(t, 1, calculate("3-4/2"))
}

func TestCalc2(t *testing.T) {
	assert.Equal(t, 42, calculateV2("42"))
	assert.Equal(t, 1, calculateV2("3-4/2"))
}
