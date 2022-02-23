package week1

import (
	"fmt"
	"testing"
)

func TestPalindromeString(t *testing.T) {
	fmt.Println(isPalindromeStr(""))
	fmt.Println(isPalindromeStr("  a    "))
	fmt.Println(isPalindromeStr("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindromeStr("race a car"))
	fmt.Println(isPalindromeStr("@CaTnAc#"))
}

func TestPalindromeString2(t *testing.T) {
	fmt.Println(isPalindromeStr2(""))
	fmt.Println(isPalindromeStr2("  a    "))
	fmt.Println(isPalindromeStr2("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindromeStr2("race a car"))
	fmt.Println(isPalindromeStr2("@CaTnAc#"))
}
