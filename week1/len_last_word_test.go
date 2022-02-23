package week1

import (
	"fmt"
	"testing"
)

func TestIsLetter(t *testing.T) {
	fmt.Println(lengthOfLastWord( "Hello World"))
	fmt.Println(lengthOfLastWord( "  Hello 56World  "))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
}
