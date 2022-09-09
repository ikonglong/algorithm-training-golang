package week3

// 删除连续重复出现 3 次的字符
func removeRepeatedChar(s string) string {
	// 用一个栈记录还不能消除的字符和它连续出现的次数
	stack := make([]*charCnt, 0, len(s))
	// 遍历字符串，假设只包含英文字符
	for _, r := range s {
		if len(stack) == 0 { // stack is empty
			stack = append(stack, &charCnt{r: r, cnt: 1})
		} else {
			// 当前字符跟栈顶字符是否相同。
			// 若相同，计数加 1 后判断是否等于 3。若等于，弹出栈顶元素。若小于，更新栈顶元素计数。
			// 若不同，压栈
			topIdx := len(stack) - 1
			top := stack[topIdx]
			if r == top.r {
				if top.cnt+1 == 3 {
					// 出栈
					stack[topIdx] = nil // Erase the top element
					stack = stack[:topIdx]
				} else {
					top.cnt = top.cnt + 1
				}
			} else {
				stack = append(stack, &charCnt{r: r, cnt: 1})
			}
		}
	}

	runes := make([]rune, 0, len(s))
	for _, c := range stack {
		for i := 1; i <= c.cnt; i++ {
			runes = append(runes, c.r)
		}
	}
	return string(runes)
}

type charCnt struct {
	r   rune
	cnt int
}
