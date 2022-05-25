package week2

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//https://leetcode-cn.com/problems/palindrome-linked-list/

// 算法：
// 1. 先通过快慢指针找到中间节点，并把慢指针扫过的节点放入 map，key 为节点索引（从 0 开始），val 为节点本身。
// 2. 将 slow 指针移至对称点右侧的第一个节点上，为比较做准备。根据 fast 是否为 nil 判断节点数是奇或偶。若为偶数，slow 此时已经到达目标位置，无需移动。
// 3. 从 slow 开始，遍历剩余节点。每个节点的对称节点是 map【--idx]。逐对比较。
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return true
	}
	if head.Next.Next == nil {
		return head.Val == head.Next.Val
	}
	idx := 0
	slow, fast := head, head
	aMap := make(map[int]int)
	for fast != nil && fast.Next != nil {
		aMap[idx] = slow.Val
		slow = slow.Next
		fast = fast.Next.Next
		idx++
	}
	//if idx % 2 != 0 {
	//	slow = slow.Next
	//}
	if fast != nil {
		slow = slow.Next
	}
	for slow != nil {
		idx--
		if v, found := aMap[idx]; found && v == slow.Val {
			slow = slow.Next
		} else {
			return false
		}
	}
	return true
}

// 算法：
// 1. 通过快慢指针，找到中间节点。查找结束时，若节点数为偶数，则慢指针指向后半部分的首节点；
//    若为奇数，则指向真正的中间节点，将这个节点看作前半部分。
// 2. 反转后半部分。因为遍历链表时无法从后向前遍历
// 3. 判断是否是回文，记录结果。前后两个子链表逐对比较。
// 4. 再次反转后半部分
func isPalindromeV2(head *ListNode) bool {
	// 空链表处理

	// 查找中间节点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil { // 节点数为偶数
		slow = slow.Next // 将中间节点看作前半部分
	}

	// 反转后半部分
	reversed := reverse(slow)

	// 逐对比较，判断是否是回文。没有对称节点的节点，其对称节点是它自己，即认为值相等。
	isPalindrome := true
	p1, p2 := head, reversed
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			isPalindrome = false
			break
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	// 再次反转
	reverse(reversed)

	return isPalindrome
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur, following := head, head.Next
	// Fixed Bug 循环条件 cur.Next != nil，报错：panic: runtime error: invalid memory address or nil pointer dereference，
	// 对应代码：tmp := following.Next。原因分析：
	// TODO
	for cur.Next != nil {
		for following != nil {
			tmp := following.Next
			following.Next = cur
			// Fixed Bug 只有头节点才需要置空
			//cur.Next = nil
			cur = following
			following = tmp
		}
	}
	head.Next = nil
	return cur
}

// --- Recursion-based solution begin

func isPalindrome_recursion(head *ListNode) bool {
	s := &Solution{
		forward: head,
	}
	return s.cmp(head)
}

type Solution struct {
	forward *ListNode // 正向走的指针的
}

// Fixed Bug 使用值接收者结果就不对，使用指针接收者就对了
//func (s Solution) cmp(head *ListNode) bool {
func (s *Solution) cmp(head *ListNode) bool {
	if head != nil {
		if !s.cmp(head.Next) { // 如果子列表非回文，直接返回 false
			return false
		}
		// 如果子列表是回文，比较逆向指针 head 和正向指针 forward，此时它们的位置刚好相对
		if head.Val != s.forward.Val {
			return false
		}
		// 如果为 true，就需要移动正向指针。
		// 这句为什么不放在 if 块之后呢？因为
		// 1) 最后一次递归调用是 cmp(head=tail.Next)，这次调用直接执行完成，无需压栈
		// 2) 之前的每次调用都需要压栈，栈帧中的 head 不为 nil。所有栈帧中的 head 构成逆向列表
		s.forward = s.forward.Next
	}
	// 如果前面的检查通过，则 head 表示的列表是回文。或者，
	// 递归到 tail.Next，此时 head 为 nil，没有要比较的节点，可认为空列表是回文的。
	return true
}

// --- Recursion-based solution end
