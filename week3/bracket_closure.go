package week3

func isValid(s string) bool {
	stack := NewStack(len(s))
	for _, r := range s {
		if r == '(' || r == '[' || r == '{' {
			stack.push(r)
		} else {
			if stack.isEmpty() { // Case: input ']'
				return false
			}
			top := stack.pop().(rune)
			if (r == ')' && top != '(') || (r == ']' && top != '[') || (r == '}' && top != '{') {
				return false
			}
		}
	}

	// Fixed Bug: case input '['
	//return true
	return stack.isEmpty()
}
