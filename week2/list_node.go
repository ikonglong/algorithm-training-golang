package week2

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	var b strings.Builder
	b.Grow(100)
	fmt.Fprint(&b, "[")
	for p := n; p != nil; {
		fmt.Fprint(&b, p.Val)
		if p.Next != nil {
			fmt.Fprint(&b, ",")
		}
		p = p.Next
	}
	fmt.Fprint(&b, "]")
	return b.String()
}

func (n *ListNode) Println() {
	fmt.Println(n.String())
}
