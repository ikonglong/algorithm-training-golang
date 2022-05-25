package week2

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// https://leetcode-cn.com/problems/middle-of-the-linked-list/
func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 一个结点
	if head.Next == nil {
		return head
	}
	// 两个结点
	if head.Next.Next == nil {
		return head.Next
	}

	// 前面两种不用单独处理
	slow := head
	fast := slow
	// FIXED BUG 忘记条件 fast != nil 了
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
