package week3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdvancedCalc(t *testing.T) {
	assert.Equal(t, 0, calculate2("0"))
	assert.Equal(t, 2, calculate2("1+1"))
	assert.Equal(t, 4, calculate2("6-4/2"))
	assert.Equal(t, 21, calculate2("2*(5+5*2)/3+(6/2+8)"))
	assert.Equal(t, -12, calculate2("(2+6*3+5-(3*14/7+2)*5)+3"))
	assert.Equal(t, 7, calculate2(" 3 +  (1*5+2)-  3"))
	assert.Equal(t, -2, calculate2(" (8 -  (1*5+2))-  3"))
}
