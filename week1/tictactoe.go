package week1

// 井字游戏（圈叉游戏）
// https://leetcode-cn.com/problems/tic-tac-toe-lcci/

func tictactoe(board []string) string {
	len := len(board)
	rowSum := 0
	colSum := 0
	diagonal := 0
	backDiagonal := 0
	hasBlank := false

	for i := 0; i < len; i++ {
		rowSum, colSum = 0, 0
		for j := 0; j < len; j++ {
			rowSum = rowSum + int(board[i][j])
			// 把 j 用作行索引，i 保持不变，就求得列之和。很巧妙
			colSum = colSum + int(board[j][i])
			if board[i][j] == ' ' {
				hasBlank = true
			}
		}
		if rowSum == 'X'*len || colSum == 'X'*len {
			return "X"
		}
		if rowSum == 'O'*len || colSum == 'O'*len {
			return "O"
		}

		// 正对角线上的元素为 board[i][i]
		diagonal = diagonal + int(board[i][i])
		// 反对角线上的元素为 board[i][len-1-i]
		backDiagonal = backDiagonal + int(board[i][len-1-i])
	}

	if diagonal == 'X'*len || backDiagonal == 'X'*len {
		return "X"
	}
	if diagonal == 'O'*len || backDiagonal == 'O'*len {
		return "O"
	}
	if hasBlank {
		return "Pending"
	} else {
		return "Draw"
	}
}

func tictactoe_bug(board []string) string {
	hasBlank := false
	cTotals := make(map[int]int)
	left, lIdx := 0, 0
	right, rIdx := 0, len(board[0])-1
	n := len(board)
	for i := 0; i < n; i++ {
		rTotal := 0
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == ' ' {
				hasBlank = true
			}
			e := toInt(board[i][j])
			rTotal += e
			//cTotals[j] = cTotals[j] + e
			if v, found := cTotals[j]; found {
				cTotals[j] = v + e
			} else {
				cTotals[j] = 0 + e
			}
			if j == lIdx {
				left += e
			}
			if j == rIdx {
				right += e
			}
		}
		lIdx++
		rIdx--
		if w := winParty(rTotal, n); w != "" {
			return w
		}
	}

	for _, v := range cTotals {
		if w := winParty(v, n); w != "" {
			return w
		}
	}
	if w := winParty(left, n); w != "" {
		return w
	}
	if w := winParty(right, n); w != "" {
		return w
	}

	if hasBlank {
		return "Pending"
	} else {
		return "Draw"
	}
}

func winParty(t int, n int) string {
	if t == 1*n {
		return "O"
	} else if t == 2*n {
		return "X"
	}
	return ""
}

func toInt(r byte) int {
	if r == ' ' {
		return 0
	} else if r == 'O' {
		return 1
	} else { // 'X'
		return 2
	}
}
