package week3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStackQueV2(t *testing.T) {
	q := Constructor()
	assert.Equal(t, 0, q.n)

	q.AppendTail(1)
	assert.Equal(t, 1, q.n)
	q.AppendTail(2)
	q.AppendTail(3)
	assert.Equal(t, 3, q.n)

	assert.Equal(t, 1, q.DeleteHead())
	assert.Equal(t, 2, q.n)
	assert.Equal(t, 2, q.DeleteHead())
	assert.Equal(t, 1, q.n)
	assert.Equal(t, 3, q.DeleteHead())
	assert.Equal(t, 0, q.n)
	assert.Equal(t, true, q.isEmpty())
}
