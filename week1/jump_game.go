package week1

// https://leetcode-cn.com/problems/jump-game/
// 考虑一个输入 [4, 0, 0, 0, 4]
func canJump(nums []int) bool {
	len := len(nums)
	// 当前元素距离最后一个元素的步数
	steps := 1
	// 从倒数第 2 个开始遍历。令结束元素指向倒数第一个元素
	for i := len - 2; i >= 0; i-- {
		if nums[i] >= steps {
			// 如果从当前元素可以跳到结束元素，令结束元素指向当前元素，
			// 在下一轮迭代中查看其前一个元素是否可以跳到它
			steps = 1
		} else {
			// 如果当前元素无法跳到结束元素，则不修改结束元素指针。
			// 这意味着下一个被迭代的元素到结束元素的距离为 steps + 1
			steps++
		}
	}
	if steps > 1 { // 如果 steps 大于 1，则表明可从第一个位置跳到最后一个位置。
		return false
	} else {
		return true
	}
}

func canJump2(nums []int) bool {
	if nums == nil || len(nums) == 0 || len(nums) == 1 {
		return true
	}

	// 想象你是行走在格子上的小人。格子里的数字代表「能量」，你需要能量才能继续行走。
	// 初始时当前能量为第一个格子的能量。每走进一个格子，你检查当前能量是否小于格子能量。
	// 如果是，就放弃当前能量，吸收格子能量，继续向前。如果否，能量不变，继续向前。
	i := 0
	cur := nums[0]
	// 当前能量大于 0 才能向前走
	for i = 1; cur > 0 && i < len(nums); i++ {
		cur-- // 每向前一步就消耗能量
		if nums[i] > cur {
			// 对于走剩余路程，这两份能量只能用一份，因此不能累加能量
			cur = nums[i]
		}
	}
	// 判断是否走到走后一个元素
	return i == len(nums)
}

// TODO 动态规划实现
