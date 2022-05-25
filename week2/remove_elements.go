package week2

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	// 处理空链表
	if head == nil {
		return nil
	}

	dummyHead := ListNode{
		Next: head,
	}
	prev := &dummyHead
	p := head
	for p != nil {
		if p.Val == val {
			prev.Next = p.Next
			p.Next = nil
			p = prev.Next
		} else {
			prev = p
			p = p.Next
		}
	}
	return dummyHead.Next
}
