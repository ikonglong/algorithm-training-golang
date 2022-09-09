package week3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack(3)
	s.push(1)
	s.push(2)
	s.push(3)
	assert.Equal(t, 3, s.n)
	assert.Equal(t, 3, s.pop().(int))
	assert.Equal(t, 2, s.n)
	assert.Equal(t, 2, s.pop().(int))
	assert.Equal(t, 1, s.n)
	assert.Equal(t, 1, s.pop().(int))
	assert.Equal(t, 0, s.n)
	assert.Equal(t, true, s.isEmpty())
}
