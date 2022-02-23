package week1

func lengthOfLastWord(s string) int {
	// 从尾部开始遍历字符串的字节，遇到第一个是字母的字节，就开始计数。
	// 在此之后，遇到第一个不是字母的字节，结束计数
	c := 0
	for i := len(s) - 1; i >= 0; i-- {
		if c >= 0 && isLetter(rune(s[i])) {
			c++
		} else if c > 0 && !isLetter(rune(s[i])) {
			break
		}
	}
	return c
}

func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}


