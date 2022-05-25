package week1

import "sort"

func isStraight(nums []int) bool {
	if nums == nil {
		return false
	}
	sort.Sort(sort.IntSlice(nums))
	zeroCount := 0
	gap := 0
	lastNonzeroIdx := -1
	for i, e := range nums {
		if e == 0 {
			zeroCount++
		} else {
			if lastNonzeroIdx >= 0 {
				// 如果前后（（位置不一定相邻））两个非零元素相同，则不是顺子
				diff := e - nums[lastNonzeroIdx]
				if diff == 0 {
					return false
				}
				// 如果相邻的两个元素不同，则累积缺口
				gap = gap + (diff - 1)
			}
			lastNonzeroIdx = i
		}
	}
	return gap <= zeroCount
}

func isStraightWithBug(nums []int) bool {
	if nums == nil {
		return false
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return false
		}
	}
	return true
}
