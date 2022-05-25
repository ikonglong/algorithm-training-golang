package week2

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head
	// Fixed Bug: slow, fast 初始值刚好满足了循环退出条件，循环压根不会执行
	//for (fast != nil && fast.Next != nil) || slow != fast {
	//	slow = slow.Next
	//	fast = fast.Next.Next
	//}
	// Fixed Bug:
	// for ok := true; ok; ok = (fast != nil && fast.Next != nil) && slow != fast {
	for ok := true; ok; ok = (fast != nil && fast.Next != nil) && slow != fast {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast == nil || fast.Next == nil {
		return false
	}
	return true
}
