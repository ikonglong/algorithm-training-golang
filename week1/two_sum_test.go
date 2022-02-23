package week1

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	fmt.Printf("%v\n", twoSum(nil, 2))
	fmt.Printf("%v\n", twoSum([]int {0, 4, -5, 7, 9}, 2))
}
