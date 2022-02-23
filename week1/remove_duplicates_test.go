package week1

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	fmt.Println(removeDuplicates(nil))
	fmt.Println(removeDuplicates([]int{}))
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums), nums)
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums), nums)
}
