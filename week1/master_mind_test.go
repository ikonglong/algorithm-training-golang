package week1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMasterMind(t *testing.T) {
	assert.Equal(t, []int{4, 0}, masterMind("RGBY", "RGBY"))
	assert.Equal(t, []int{0, 0}, masterMind("RGGY", "BBBB"))
	assert.Equal(t, []int{1, 0}, masterMind("GGGG", "BBGB"))
	assert.Equal(t, []int{1, 1}, masterMind("RGBY", "GGRR"))
	assert.Equal(t, []int{1, 2}, masterMind("RGGY", "GGRR"))
}

func TestMasterMindV2(t *testing.T) {
	//assert.Equal(t, []int{4, 0}, masterMind("RGBY", "RGBY"))
	//assert.Equal(t, []int{0, 0}, masterMind("RGGY", "BBBB"))
	//assert.Equal(t, []int{1, 0}, masterMind("GGGG", "BBGB"))
	//assert.Equal(t, []int{1, 1}, masterMind("RGBY", "GGRR"))
	assert.Equal(t, []int{1, 2}, masterMindV2("RGGY", "GGRR"))
	assert.Equal(t, []int{1, 2}, masterMindV2("RGGY", "GGRR"))
}
