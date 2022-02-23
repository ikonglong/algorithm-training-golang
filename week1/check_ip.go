package week1

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// 题目：
// 给定一个字符串表示的 IP 地址，例如 "123.92.2.34"，判断其是否合法。合法 IP 地址的规则如下：
// a. 除了空格、数字和 . 之外，不得包含其他字符
// b. IP 地址由四个数字构成，由 . 分隔，每个 . 分隔的数字大小在 0~255 之间
// c. 数字前后可以有空格，但中间不能有空格。例如，"123 . 92. 2 . 34" 合法，"12 3.9 2.2.34"非法。
//
// 当然，这个问题还可以继续添加一些规则，让题目变得更加复杂。比如每个数字不能有前导 0，但可以为 0。例如，"021.3.02.34" 非法，"0.2.0.33" 合法。

func checkIp(ip string) bool {
	if ip == "" {
		return false
	}

	// split ip into segments
	segments := strings.Split(ip, ".")
	// check segment
	for _, seg := range segments {
		if !checkSegment(seg) {
			return false
		}
	}
	return true
}

func checkSegment(seg string) bool {
	if seg == "" {
		return false
	}

	// Ignore leading whitespaces
	var i = 0
	for ; i < len(seg); {
		r, size := utf8.DecodeRuneInString(seg[i:])
		if r != ' ' {
			break;
		}
		i += size
	}

	var num = 0
	for ; i < len(seg); {
		r, size := utf8.DecodeRuneInString(seg[i:])
		if r == ' ' {
			break // 遇到空格就认为 seg 的有效字符已处理完成
		} else if unicode.IsDigit(r) {
			num += num * 10 + int(r - '0')
		} else {
			return false // 如果非数字
		}
		i += size
	}
	// 判断 seg 表示的数字
	if num > 255 {
		return false
	}

	// Check trailing chars
	for ; i < len(seg); {
		r, size := utf8.DecodeRuneInString(seg[i:])
		if r != ' ' {
			return false // 只要非空格，当前 seg 就非法
		}
		i += size
	}

	return true
}
