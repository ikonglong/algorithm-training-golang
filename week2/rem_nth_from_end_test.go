package week2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemNthFromEnd(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	// 因为递归开始于 preHead，preHead 是倒数第 4 个节点（不是期望的状况），
	// 所以无法遍历到其前驱节点，因此无法删除 preHead，这碰巧产生了期望的结果。
	rs := removeNthFromEnd(head, 4)
	rs.Print()
	assert.Equal(t, "[1,2,3]", rs.String())
}
