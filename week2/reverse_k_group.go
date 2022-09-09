package week2

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k <= 1 {
		return head
	}
	p := head
	cnt := 0
	p, kHead := head, head
	var kTail *ListNode = nil
	preNewHead := &ListNode{}
	preKTail := preNewHead
	for p != nil {
		cnt++
		// Fixed Bug: reverseK 会修改 p.Next
		next := p.Next
		if cnt%k == 0 {
			kTail = p
			preKTail.Next = reverseK(kHead, kTail)
			preKTail = kHead

			kHead = next
			kTail = nil
		}
		p = next
	}
	if cnt%k != 0 {
		preKTail.Next = kHead
	}
	return preNewHead.Next
}

func reverseK(head *ListNode, tail *ListNode) *ListNode {
	//pre := &ListNode{}
	var pre *ListNode = nil
	cur, end := head, tail.Next
	//for cur != tail.Next {
	for cur != end {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return tail
}
