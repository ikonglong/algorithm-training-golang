package week3

import "log"

func calculate(s string) int {
	numStack := NewStack(8)
	opStack := NewStack(8)

	curNum := 0
	for _, r := range s {
		if r == ' ' {
			continue
		} else if isDigit(r) {
			curNum += curNum*10 + int(r-'0')
		} else {
			numStack.push(curNum)
			curNum = 0

			op, found := operators[r]
			if !found {
				log.Panicf("Bad operator: %v", r)
			}
			// 如果待处理的运算符优先于栈顶运算符，则直接添加至栈顶。
			// 如果不优先于后者，则执行栈顶运算符表示的计算。重复，直到没有计算可做或待处理运算符优先于栈顶运算符。
			for !opStack.isEmpty() && !op.prior(opStack.peek().(Operator)) {
				rs := fetchAndCalc(numStack, opStack)
				numStack.push(rs)
			}
			opStack.push(op)
		}
	}
	// Fixed Bug 前面的 for 循环中遇到运算符才能把数字放入栈中。对于最后一个数字，循环结束时还没有放入栈中。
	// 更好的写法是，在循环中嵌入一个子循环
	numStack.push(curNum)

	// 注意，最后运算符栈只可能剩下优先级相同的运算符。
	// 例如，计算 3-4/2
	for !opStack.isEmpty() {
		rs := fetchAndCalc(numStack, opStack)
		numStack.push(rs)
	}

	return numStack.pop().(int)
}

func calculateV2(s string) int {
	numStack := NewStack(8)
	opStack := NewStack(8)

	for i := 0; i < len(s); {
		r := s[i]
		if r == ' ' {
			i++
			continue
		}
		num := 0
		if isDigit(rune(r)) {
			for i < len(s) && isDigit(rune(s[i])) {
				num = num*10 + int(s[i]-'0')
				i++
			}
			numStack.push(num)
		} else {
			op := operators[rune(r)]
			for !opStack.isEmpty() && !op.prior(opStack.peek().(Operator)) {
				rs := fetchAndCalc(numStack, opStack)
				numStack.push(rs)
			}
			opStack.push(op)
			i++
		}
	}

	for !opStack.isEmpty() {
		rs := fetchAndCalc(numStack, opStack)
		numStack.push(rs)
	}

	return numStack.pop().(int)
}

func fetchAndCalc(numStack *Stack, opStack *Stack) int {
	num2, num1 := numStack.pop().(int), numStack.pop().(int)
	op := opStack.pop().(Operator)
	switch op.symbol {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	default:
		log.Fatalf("Bad operator: %v", op.symbol)
		return -1
	}
}

type Operator struct {
	symbol   string
	priority int
}

func (op Operator) prior(op2 Operator) bool {
	return op.priority > op2.priority
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

var operators map[rune]Operator = map[rune]Operator{
	rune('+'): {
		symbol:   "+",
		priority: 1,
	},
	rune('-'): {
		symbol:   "-",
		priority: 1,
	},
	rune('*'): {
		symbol:   "*",
		priority: 2,
	},
	rune('/'): {
		symbol:   "/",
		priority: 2,
	},
}
