package week3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestXxx(t *testing.T) {
	assert.False(t, isValid("["))
	assert.False(t, isValid("]"))
	assert.True(t, isValid("[]"))
	assert.True(t, isValid("[()]"))
}
