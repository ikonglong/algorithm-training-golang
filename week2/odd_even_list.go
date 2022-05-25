package week2

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	evenHead := head.Next
	oddTail, evenTail, p := head, head.Next, head.Next.Next
	odd := true
	for p != nil {
		if odd {
			oddTail.Next = p
			oddTail = p
		} else {
			evenTail.Next = p
			evenTail = p
		}
		p = p.Next
		odd = !odd
	}
	evenTail.Next = nil
	oddTail.Next = evenHead
	return head
}
