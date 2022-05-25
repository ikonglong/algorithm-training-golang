package week2

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return head
	}
	preHead := &ListNode{
		Val:  0,
		Next: head,
	}
	remKthFromEnd(preHead, n)
	return preHead.Next
}

func remKthFromEnd(p *ListNode, k int) int {
	if p == nil {
		return 0
	}
	nextCnt := remKthFromEnd(p.Next, k)
	if nextCnt == k {
		p.Next = p.Next.Next
	}
	return 1 + nextCnt
}
