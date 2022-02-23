package week1

import (
	"fmt"
	"testing"
)

func TestIsPalindromeNum(t *testing.T) {
	fmt.Println(isPalindrome(0))
	fmt.Println(isPalindrome(101))
	fmt.Println(isPalindrome(-1015101))
}
