package week1

import (
	"fmt"
	"testing"
)

func TestReverseLeftWords(t *testing.T) {
	fmt.Println(reverseLeftWords("abcdefg", 2))
	fmt.Println(reverseLeftWords("abcdefg", 7))
	fmt.Println(reverseLeftWords("abcdefg", 9))
	fmt.Println(reverseLeftWords("abcdefg", 0))
	fmt.Println(reverseLeftWords("lrloseumgh", 6))
}
