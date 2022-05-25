package week2

import (
	"fmt"
	"testing"
)

func TestRemoveMiddle(t *testing.T) {
	n6 := ListNode{
		Val:  6,
		Next: nil,
	}
	n5 := ListNode{
		Val:  5,
		Next: &n6,
	}
	n4 := ListNode{
		Val:  4,
		Next: &n5,
	}
	n3 := ListNode{
		Val:  3,
		Next: &n4,
	}
	n2 := ListNode{
		Val:  2,
		Next: &n3,
	}
	n1 := ListNode{
		Val:  1,
		Next: &n2,
	}
	fmt.Printf("%v\n", middleNode(&n1))
}
