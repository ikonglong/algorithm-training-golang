package week2

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	slow, fast := head, head
	for {
		// 如果 fast.Next 或 fast.Next.Next 为 nil，就不用向前跳了。
		// 这样也可以保证，只要初始时 fast 不为 nil，循环过程中 fast 不可能为 nil。
		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast { // 环上相遇
			break
		}
	}
	ptr := head
	for ptr != slow {
		ptr = ptr.Next
		slow = slow.Next
	}
	return ptr
}
