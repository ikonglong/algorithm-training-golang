package week3

type charCntV2 struct {
	idx int
	cnt int
}

func removeDuplicates(s string) string {
	stack := make([]charCntV2, 0, len(s))
	for i, r := range s {
		n := len(stack)
		if n == 0 { // case: stack is empty
			stack = append(stack, charCntV2{
				idx: i,
				cnt: 1,
			})
			continue
		}
		topRune := rune(s[stack[n-1].idx])
		if r != topRune {
			stack = append(stack, charCntV2{
				idx: i,
			})
		} else {
			stack = stack[:n-1]
		}
	}

	runes := make([]rune, 0, len(stack))
	for _, e := range stack {
		runes = append(runes, rune(s[e.idx]))
	}
	return string(runes)
}

func removeDuplicatesV2(s string) string {
	stack := make([]charCntV2, 0, len(s))
	for i, r := range s {
		n := len(stack)
		// Fixed Bug
		// 将判断「栈是否为空」和「当前字符是否跟栈顶元素相同」合并在一个 if 判断中很容易
		// 引入 bug。例如在检查是否为空前就 peek。
		if n == 0 || r != rune(s[stack[n-1].idx]) {
			stack = append(stack, charCntV2{
				idx: i,
			})
		} else {
			stack = stack[:n-1]
		}
	}

	runes := make([]rune, 0, len(stack))
	for _, e := range stack {
		runes = append(runes, rune(s[e.idx]))
	}
	return string(runes)
}
