package week2

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA, pB := headA, headB

	// Fixed Bug: 低级错误！把退出循环的条件用作循环条件了。
	//for pA == pB || (pA == nil && pB == nil) {
	// Fixed Bug: 低级错误！pA, pB 为空时，将 pA, pB 指向 headB, head 后，紧接着的操作应该是比较，而非移动指针。
	// 因为处理空值就是在移动指针到下一个节点。
	//	if pA == nil {
	//		pA = headB
	//	}
	//	if pB == nil {
	//		pB = headA
	//	}
	//	pA = pA.Next
	//	pB = pA.Next
	//}

	// 两个同为 nil 时，pA == pB 为真
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}
