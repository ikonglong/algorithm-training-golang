package week3

// 用两个栈实现队列
// https://leetcode.cn/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/

// V1 实现思路：
// 一个栈用作尾部栈，另一个用作头部栈。
// 入队时，压入尾部栈。出队时，先从尾部栈转储至头部栈，接着从头部栈 pop，然后再从头部栈转储至尾部栈。

// TODO

// ~~~~~~~~~~~~~~~~~~

// 一个栈用作队列头部，一个栈用作队列尾部。
// 入队时，直接压入尾部栈。
// 出队时，如果头部栈不为空，直接从头部栈 pop；如果为空，先从尾部栈转储至头部栈，再从头部栈 pop。
type CQueueV2 struct {
	headStack *Stack // 用作队列头部
	tailStack *Stack // 用作队列尾部
	n         int    // 元素个数
}

func Constructor() CQueueV2 {
	return CQueueV2{
		headStack: NewStack(16),
		tailStack: NewStack(16),
		n:         0,
	}
}

func (q *CQueueV2) AppendTail(value int) {
	q.tailStack.push(value)
	q.n++
}

func (q *CQueueV2) DeleteHead() int {
	if q.isEmpty() {
		panic("this que is empty")
		return -1
	}
	q.n--
	if !q.headStack.isEmpty() {
		return q.headStack.pop().(int)
	}
	// headStack 必须为空，否则队列中元素顺序就乱了
	q.dumpTailToHead()
	return q.headStack.pop().(int)
}

func (q *CQueueV2) dumpTailToHead() {
	for !q.tailStack.isEmpty() {
		q.headStack.push(q.tailStack.pop())
	}
}

func (q *CQueueV2) isEmpty() bool {
	return q.n == 0
}
