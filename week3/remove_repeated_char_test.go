package week3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveRepeatedChar(t *testing.T) {
	s := removeRepeatedChar("abbaaabac")
	assert.Equal(t, "aac", s)
}
