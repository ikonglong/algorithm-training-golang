package week2

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var prev *ListNode = nil
	cur := head
	for cur != nil {
		if prev != nil && prev.Val == cur.Val {
			prev.Next = cur.Next // 删除当前节点
			cur = cur.Next       // 当前节点右移一位。注意 prev 指针不能动
		} else {
			prev = cur
			cur = cur.Next
		}
	}
	return head
}
