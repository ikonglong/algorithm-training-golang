package week1

func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	// 两个指针 i, j
	// 指针 i 指向从头部开始的不重复部分的结束位置，
	// 指针 j 从 i + 1 开始向后扫描，
	i := 0
	j := i + 1
	for i < len(nums) && j < len(nums) {
		if nums[i] == nums[j] {
			j++
		} else {
			// j >= i+1 must be true
			if j > i+1 {
				nums[i+1] = nums[j]
			}
			i++
			j++
		}
	}
	return i + 1
}
