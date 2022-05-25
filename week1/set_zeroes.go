package week1

// https://leetcode-cn.com/problems/zero-matrix-lcci/
func setZeroes(matrix [][]int) {
	if matrix == nil {
		return
	}
	// 这两个 map 也可以替换为两个 bool 类型的标记数组，
	// 其长度分别为行长度、列长度。
	rowSet := make(map[int]bool)
	colSet := make(map[int]bool)
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			if matrix[r][c] == 0 {
				rowSet[r] = true
				colSet[c] = true
			}
		}
	}
	for r := range rowSet {
		for i := 0; i < len(matrix[r]); i++ {
			matrix[r][i] = 0
		}
	}
	for c := range colSet {
		for r := 0; r < len(matrix); r++ {
			matrix[r][c] = 0
		}
	}
}

func setZeroesV2_bug(matrix [][]int) {
	if matrix == nil {
		return
	}

	// 先用两个标记变量表示第一行和第一列是否包含 0，再用第一行和第一列表示剩余列和行中是否有 0。
	zeroInRow0 := false
	zeroInCol0 := false
	for r := 0; r < len(matrix); r++ {
		if matrix[r][0] == 0 {
			zeroInCol0 = true
		}
		for c := 0; c < len(matrix[r]); c++ {
			if r == 0 {
				if matrix[r][c] == 0 {
					zeroInRow0 = true
				}
			} else if matrix[r][c] == 0 {
				matrix[r][0], matrix[0][c] = 0, 0
			}
		}
	}

	// --- start: code snippet with bug
	//for r := len(matrix) - 1; r >= 0; r-- {
	//	// 这里有 bug。如果 m[0][0] 为 0，那么内循环第一次执行先据此将 m[r][0] 设置为 0，
	//	// 然后根据 m[r][0] 为 0，将剩余每一列都设置为 0
	//	for c := 0; c < len(matrix[r]); c++ {
	//		if matrix[0][c] == 0 || matrix[r][0] == 0 {
	//			matrix[r][c] = 0
	//		}
	//	}
	//}
	// --- end: code snippet with bug

	for r := len(matrix) - 1; r >= 1; r-- {
		for c := 1; c < len(matrix[r]); c++ {
			if matrix[0][c] == 0 || matrix[r][0] == 0 {
				matrix[r][c] = 0
			}
		}
	}

	if zeroInRow0 {
		for c := 0; c < len(matrix[0]); c++ {
			matrix[0][c] = 0
		}
	}
	if zeroInCol0 {
		for r := 0; r < len(matrix); r++ {
			matrix[r][0] = 0
		}
	}
}

func setZeroesV3(matrix [][]int) {
	if matrix == nil {
		return
	}

	// 先用一个变量标记第一列是否包含 0。这样此后就可以将每行中的 0 投影到该行的第一列。
	// 每一列中的 0 投影到该列的第一行。这样第一行第一列为 0 就无歧义了，仅表示第一行有 0。
	zeroInCol0 := false
	for r := 0; r < len(matrix); r++ {
		if matrix[r][0] == 0 {
			zeroInCol0 = true
		}
		for c := 1; c < len(matrix[r]); c++ {
			if matrix[r][c] == 0 {
				matrix[0][c], matrix[r][0] = 0, 0
			}
		}
	}

	// 因为第一列和第一行中的 0 含义特殊，表示对应行和列有 0，而非数字 0 本身。
	// 所以在置为 0 的阶段要防止第一行先于其他行被设置为 0，得从最后一行向前处理。
	for r := len(matrix) - 1; r >= 0; r-- {
		for c := 1; c < len(matrix[r]); c++ {
			// 检查当前单元格所在的列 和 行是否包含 0
			if matrix[0][c] == 0 || matrix[r][0] == 0 {
				matrix[r][c] = 0
			}
		}
		if zeroInCol0 {
			matrix[r][0] = 0
		}
	}
}
