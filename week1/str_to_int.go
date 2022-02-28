package week1

import "math"

func strToInt(str string) int {
	if str == ""  {
		return 0
	}

	// 忽略前导空格
	i := 0
	for ; i < len(str); i++ {
		if str[i] != ' ' {
			break
		}
	}
	if i == len(str) { // 空字符串
		return 0
	}

	// 若第一个非空字符不是数字、+、-，则返回 0。
	// 之后的字符：
	// 若为 +，则保持不变；若为 -，则乘以 -1；
	// 若为数字，则先将结果左移一位（x10），再加上该数字；
	// 若遇到非数字字符，结束循环
	rs := 0
	sign := 1
	for j := i; j < len(str); j++ {
		r := str[j]
		if j == i && r == '+' {
			// do nothing
		} else if j == i && r == '-' {
			sign = -1
		} else if r >= '0' && r <= '9' {
			rs = rs * 10 + int(r - '0') * sign
			if rs > math.MaxInt32  {
				rs = math.MaxInt32
				break
			} else if rs < math.MinInt32 {
				rs = math.MinInt32
				break
			}
		} else {
			break
		}
	}

	return rs
}
