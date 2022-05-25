package week1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsStraight(t *testing.T) {
	assert.True(t, isStraight([]int{1, 2, 3, 4, 5}))
	assert.True(t, isStraight([]int{0, 0, 5, 2, 0}))
	assert.False(t, isStraight([]int{0, 0, 1, 3, 13}))
	// 这个 case 虽然未通过，但是扑克牌中木有负数。
	assert.False(t, isStraight([]int{0, 0, 1, 3, -5}))
	assert.False(t, isStraight([]int{0, 0, 1, 1, 1}))
	assert.False(t, isStraight([]int{0, 0, 2, 2, 5}))
}
