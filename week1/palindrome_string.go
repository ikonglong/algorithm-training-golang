package week1

import "unicode"

// 验证是否是回文字符串时，可以忽略其中的标点符号和空格
func isPalindromeStr(s string) bool {
	if s == "" {
		return true
	}

	// 用两个指针 l, r 分别从字符串的两端开始，向另一端逐个字符扫描，并对比被扫描的字符。
	// 遇到非字母和数字的字符，跳过，不对比
	for l, r := 0, len(s)-1; l < r; {
		if !isLetterOrDigit(s[l]) {
			l++
			continue
		}
		if !isLetterOrDigit(s[r]) {
			r--
			continue
		}
		if unicode.ToLower(rune(s[l])) == unicode.ToLower(rune(s[r])) {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}

func isPalindromeStr2(s string) bool {
	if s == "" {
		return true
	}

	l, r := 0, len(s)-1
	for l < r {
		for !isLetterOrDigit(s[l]) {
			l++
		}
		for !isLetterOrDigit(s[r]) {
			r--
		}
		if unicode.ToLower(rune(s[l])) == unicode.ToLower(rune(s[r])) {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}

func isLetterOrDigit(char byte) bool {
	r := rune(char)
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
