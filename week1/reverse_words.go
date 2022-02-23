package week1

import (
	"bytes"
)

// 1. 倒序遍历字符串，记录单词左右边界的索引
// 2. 每确定一个单词的边界，就将该单词添加到表示结果的 string buffer 中。如果 buf 中已有数据，添加当前单词前，先添加一个空格。
func reverseWords(s string) string {
	if s == "" {
		return s
	}

	// Ignore trailing whitespaces
	ri := -1 // 单词右边界的索引
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			ri = i
			break
		}
	}

	if ri == -1 { // 表明 s 包含的全是空格
		return ""
	}

	var buf bytes.Buffer
	for i := ri; i >= 0; i-- {
		if s[i] == ' ' {
			if ri != -1 { // 确定了一个单词的边界 i+1, ri
				if buf.Len() != 0 {
					//buf.WriteRune(' ')
					buf.WriteByte(' ')
				}
				buf.WriteString(s[i+1 : ri+1])
				ri = -1
			}
		} else if ri == -1 {
			ri = i
		}
	}

	// 遍历结束也要触发是否有单词的判断。遍历结束时，若 ri 不为 -1，表明 s[0:ri+1] 仍是一个单词
	if ri != -1 {
		if buf.Len() != 0 {
			buf.WriteByte(' ')
		}
		buf.WriteString(s[0 : ri+1])
	}

	return buf.String()
}
