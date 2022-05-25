package week2

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2
	dummyHead := &ListNode{
		Val: 0,
	}
	tail := dummyHead
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			tail.Next = p1
			p1 = p1.Next
		} else {
			tail.Next = p2
			p2 = p2.Next
		}
		tail = tail.Next
	}
	if p1 != nil {
		tail.Next = p1
	} else {
		tail.Next = p2
	}
	return dummyHead.Next
}
