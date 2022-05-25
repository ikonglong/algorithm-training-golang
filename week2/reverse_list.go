package week2

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//https://leetcode-cn.com/problems/reverse-linked-list/solution/
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	cur := head
	var next *ListNode = nil
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

// 依然是从前往后反转。跟 reverseList 相比，不同之处在于
// reverseList 用了两个指针 prev, cur，原本 prev.next = cur，修改为 cur.next = prev；
// reverseList2 用了两个指针 cur, next, 原本 cur.next = next，修改为 next.next = cur，为了节省 next 变量，巧用 head.next 保存。
func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur := head
	for head.Next != nil {
		t := head.Next.Next
		head.Next.Next = cur
		// Fixed Bug:
		//cur.Next = nil
		cur = head.Next
		head.Next = t
	}
	// 这行代码多余，因为 head.Next 最后是 nil，刚好符合期望，巧妙！
	head.Next = nil
	return cur
}

func reverseListByRecursion(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var newHead *ListNode
	// Fixed Bug:
	// panic: runtime error: invalid memory address or nil pointer dereference [recovered]
	// panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x1136484]
	//var newHead **ListNode = nil
	doReverse2(head, &newHead)
	return newHead
}

// 反转链表 head，返回反转后的链表的尾节点
func doReverse(head *ListNode, newHeadPtr **ListNode) *ListNode {
	if head.Next == nil { // 处理单节点的链表
		*newHeadPtr = head
		return head
	}
	prev := doReverse(head.Next, newHeadPtr)
	prev.Next = head
	// 不能忘，否则链表成环
	head.Next = nil
	return head
}

func doReverse2(head *ListNode, newHeadPtr **ListNode) {
	if head.Next == nil { // 处理单节点的链表
		*newHeadPtr = head
		return
	}
	doReverse2(head.Next, newHeadPtr)
	// Fixed Bug:
	// head.Next = head
	// 让当前节点的下一个节点的 Next 指针指向它自己
	head.Next.Next = head
	// 不能忘，否则链表成环
	head.Next = nil
}

func reverseListByRecursion2(head *ListNode) *ListNode {
	// head is nil 只在初次调用时会发生，因为 head.Next is nil 时不会执行递归调用
	if head == nil || head.Next == nil {
		return head
	}
	// 最后一次调用 reverseListByRecursion2 的实参是 tail 节点，返回的也是 tail 节点。
	// 然后通过递归调用逐层返回给最外层的调用者。
	r := reverseListByRecursion2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return r
}
