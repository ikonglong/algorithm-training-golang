package week3

import "log"

func calculate2(s string) int {
	nums := NewStack(8)
	ops := NewStack(8)

	for i := 0; i < len(s); {
		r := rune(s[i])
		if r == ' ' {
			i++
			continue
		}
		if isDigit(r) {
			num := 0
			for i < len(s) && isDigit(rune(s[i])) {
				num = num*10 + int(s[i]-'0')
				i++
			}
			nums.push(num)
		} else {
			if r == '(' {
				ops.push(&Operator2{
					symbol:   "(",
					priority: -1,
				})
			} else if r == ')' {
				// 遇到右括号，就重复过程：弹出一个运算符和两个操作数，计算结果再入栈，直到遇到左括号
				top := ops.peek().(*Operator2)
				for !ops.isEmpty() && top.symbol != "(" {
					nums.push(fetchAndCalc2(nums, ops))
					top = ops.peek().(*Operator2)
				}
				// 弹出左括号
				ops.pop()
			} else {
				op := operators2[r]
				// 如果栈顶运算符优先级高于等于当前运算符，就弹出一个运算符和两个操作数，运算结果再入栈。
				// 如果运算符栈为空或遇到左括号，退出循环
				calcInParentheses(nums, ops, op)
				ops.push(op)
			}
			i++
		}
	}

	for !ops.isEmpty() {
		nums.push(fetchAndCalc2(nums, ops))
	}

	return nums.pop().(int)
}

func calcInParentheses(nums *Stack, ops *Stack, op *Operator2) {
	if ops.isEmpty() {
		return
	}
	// Fixed Bug: 循环过程中 op stack 会变成空。
	for top := ops.peek().(*Operator2); top.symbol != "(" && !op.prior(top); {
		nums.push(fetchAndCalc2(nums, ops))
		if ops.isEmpty() {
			break
		}
		top = ops.peek().(*Operator2)
	}
}

func fetchAndCalc2(numStack *Stack, opStack *Stack) int {
	num2, num1 := numStack.pop().(int), numStack.pop().(int)
	op := opStack.pop().(*Operator2)
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

var operators2 map[rune]*Operator2 = map[rune]*Operator2{
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

type Operator2 struct {
	symbol   string
	priority int
}

func (op *Operator2) prior(op2 *Operator2) bool {
	return op.priority > op2.priority
}
