package week1

import (
	"fmt"
	"testing"
)

func TestReverseWords(t *testing.T) {
	fmt.Printf("reversed: %v\n", reverseWords(" a good   example  "))
	fmt.Printf("reversed: %v\n", reverseWords("a   "))
	fmt.Printf("reversed: %v\n", reverseWords("   Hello    world.   "))
	fmt.Printf("reversed: %v\n", reverseWords("     "))
}
