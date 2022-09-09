package week3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDup(t *testing.T) {
	assert.Equal(t, "ca", removeDuplicates("abbaca"))
}

func TestRemoveDupV2(t *testing.T) {
	assert.Equal(t, "ca", removeDuplicatesV2("abbaca"))
}
