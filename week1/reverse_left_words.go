package week1

func reverseLeftWords(s string, n int) string {
	if n <= 0 {
		return s
	}
	// 根据要把头部 n 个字符左旋到字符串尾部，可计算出要把尾部 len(s) - n 个字符串右旋到头部
	rn := len(s) - n
	if rn <= 0 {
		return s
	}
	//bytes := make([]byte, 0, len(s)) // len is 0, so index out of range
	bytes := make([]byte, len(s))
	for i, j := 0, rn; i < len(s); i++ {
		if i < n {
			bytes[j] = s[i]
			j++
		} else {
			bytes[i - n] = s[i]
		}
	}
	return string(bytes)
}
