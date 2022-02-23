package week1

import "strconv"

func isPalindrome(x int) bool {
	// 先把数字转换为字符串，再判断该字符串是否是回文串
	return isPalindromeNum(strconv.Itoa(x))
}

func isPalindromeNum(numStr string) bool {
	l, r := 0, len(numStr)-1
	for l < r {
		if numStr[l] == numStr[r] {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}
