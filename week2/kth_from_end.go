package week2

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/solution/
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || k <= 0 {
		return nil
	}
	var kth *ListNode = nil
	countFromEnd(head, k, &kth)
	return kth
}

func countFromEnd(p *ListNode, k int, kthPtr **ListNode) int {
	if p == nil {
		return 0
	}
	cnt := 1 + countFromEnd(p.Next, k, kthPtr)
	if cnt == k {
		*kthPtr = p
	}
	return cnt
}

// 算法：
// 1. 先遍历一次链表，得到长度 n
// 2. 倒数 k 转换为正向数第 n-k+1
// 3. 再次遍历链表，找到第 n-k+1 个
// TODO 实现

// 算法：
// 1. 构造一个栈
// 2. 遍历链表，把每个元素 push 进去
// 3. 从栈 pop 元素，计数查找第 k 个
// TODO 实现
