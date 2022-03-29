package week1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOneEditAway(t *testing.T) {
	assert.True(t, oneEditAway("", ""))
}
