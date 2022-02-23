package week1

func twoSum(nums []int, target int) []int {
	if nums == nil || len(nums) == 1 {
		return nil
	}
	res := make([]int, 0, 2)
	// 创建一个映射，用于存储当前遍历到的元素之前的元素的值到下标的映射
	valToIdx := make(map[int]int, len(nums))
	// 遍历 nums，在 valToIdx 中查找跟当前元素的和为 target 的元素
	for i, n := range nums {
		idx, ok := valToIdx[target - n]
		if ok {
			res = append(res, idx)
			res = append(res, i)
		} else {
			valToIdx[n] = i
		}
	}
	return res
}
