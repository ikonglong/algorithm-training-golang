package week2

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{
		Val: 0,
	}
	padding := &ListNode{
		Val: 0,
	}
	tail := dummyHead
	p1, p2 := l1, l2
	carry := 0
	for p1 != nil || p2 != nil {
		if p1 == nil {
			p1 = padding
		}
		if p2 == nil {
			p2 = padding
		}
		tail.Next = &ListNode{}
		tail = tail.Next
		tail.Val = (p1.Val + p2.Val + carry) % 10
		carry = (p1.Val + p2.Val + carry) / 10
		p1 = p1.Next
		p2 = p2.Next
	}
	if carry > 0 {
		tail.Next = &ListNode{
			Val: carry,
		}
	}
	return dummyHead.Next
}

func addTwoNumbers_v2(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{
		Val: 0,
	}
	tail := dummyHead
	p1, p2 := l1, l2
	carry := 0
	for p1 != nil && p2 != nil {
		tail.Next = &ListNode{}
		tail = tail.Next
		tail.Val = (p1.Val + p2.Val + carry) % 10
		carry = (p1.Val + p2.Val + carry) / 10
		p1 = p1.Next
		p2 = p2.Next
	}
	var l *ListNode
	if p1 != nil {
		l = p1
	} else if p2 != nil {
		l = p2
	}
	for p := l; p != nil; {
		tail.Next = &ListNode{}
		tail = tail.Next
		tail.Val = (p.Val + carry) % 10
		carry = (p.Val + carry) / 10
		p = p.Next
	}
	if carry > 0 {
		tail.Next = &ListNode{
			Val: carry,
		}
	}
	return dummyHead.Next
}

func addTwoNumbers_bug(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{
		Val: 0,
	}
	tail := dummyHead
	p1, p2 := l1, l2
	carry := 0
	for p1 != nil && p2 != nil {
		tail.Next = &ListNode{}
		tail = tail.Next
		tail.Val = (p1.Val + p2.Val + carry) % 10
		carry = (p1.Val + p2.Val + carry) / 10
		p1 = p1.Next
		p2 = p2.Next
	}
	// Fixed Bug: 输入 [9,9,9,9,9,9,9], [9,9,9,9]
	//if p1 != nil {
	//	p1.Val = p1.Val + carry
	//	tail.Next = p1
	//} else {
	//	p2.Val = p2.Val + 1
	//	tail.Next = p2
	//}
	var l *ListNode
	if p1 != nil {
		l = p1
	} else if p2 != nil {
		l = p2
	}
	for p := l; p != nil; {
		tail.Next = &ListNode{}
		tail = tail.Next
		tail.Val = (p.Val + carry) % 10
		carry = (p.Val + carry) / 10
		p = p.Next
	}
	return dummyHead.Next
}
