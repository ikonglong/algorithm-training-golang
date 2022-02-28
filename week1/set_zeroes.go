package week1

// https://leetcode-cn.com/problems/zero-matrix-lcci/
func setZeroes(matrix [][]int)  {
	if matrix == nil {
		return
	}
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
