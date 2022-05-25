package week1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOneEditAway(t *testing.T) {
	assert.True(t, oneEditAway("", ""))
	assert.True(t, oneEditAway("palm", "palm"))
	assert.True(t, oneEditAway("abcd", "abd"))
	assert.True(t, oneEditAway("abcd", "bcd"))
	assert.True(t, oneEditAway("abcd", "abc"))
	assert.False(t, oneEditAway("abcd", "ab"))
	assert.False(t, oneEditAway("abcd", "ad"))

	assert.True(t, oneEditAway("你在哪里呀", "你在哪里呀"))
	assert.True(t, oneEditAway("你在哪里呀", "你在哪里呀"))
	assert.True(t, oneEditAway("你在哪里呀", "在哪里呀"))
	assert.False(t, oneEditAway("你在哪里呀", "在哪里"))
}

func TestOneEditAwayV2(t *testing.T) {
	assert.True(t, oneEditAwayV2("", ""))
	assert.True(t, oneEditAwayV2("palm", "palm"))
	assert.True(t, oneEditAwayV2("abcd", "abd"))
	assert.True(t, oneEditAwayV2("abcd", "bcd"))
	assert.True(t, oneEditAwayV2("abcd", "abc"))
	assert.False(t, oneEditAwayV2("abcd", "ab"))
	assert.False(t, oneEditAwayV2("abcd", "ad"))

	//assert.True(t, oneEditAwayV2("你在哪里呀", "你在哪里呀"))
	//assert.True(t, oneEditAwayV2("你在哪里呀", "你在哪里呀"))
	//assert.True(t, oneEditAwayV2("你在哪里呀", "在哪里呀"))
	//assert.False(t, oneEditAwayV2("你在哪里呀", "在哪里"))
}
