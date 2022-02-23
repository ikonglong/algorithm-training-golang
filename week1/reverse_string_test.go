package week1

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	reverseString(nil)
	reverseString([]byte{})

	s := "hello"
	fmt.Printf("Input: %v\n", s)
	sBytes := []byte(s)
	reverseString(sBytes)
	fmt.Printf("Output: %v\n", string(sBytes))
	fmt.Printf("-------------------------\n")

	s = "ko"
	fmt.Printf("Input: %v\n", s)
	sBytes = []byte(s)
	reverseString(sBytes)
	fmt.Printf("Output: %v\n", string(sBytes))
}
