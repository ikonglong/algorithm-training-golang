package week1

// 原地旋转
// https://leetcode-cn.com/problems/rotate-image/
func rotate(matrix [][]int) {
	// TODO 文档待整理
	// 观察一个 6x6 矩阵的旋转，用 i 表示行索引，j 表示列索引。
	// 外层循环 i 取值为 [0, n/2)，内层循环 j 取值为 [i, n-1-i]。
	// 将 6x6 矩阵看作三个由单元方格组成的正方形边框嵌套组成：6x6, 4x4, 2x2。
	// 外层循环每一次迭代对一个正方形边框进行旋转，由外向内。内层循环表示为了旋转
	// 一个正方形边框，需要把边框中的方格分成几组来旋转。
	if matrix == nil {
		return
	}
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		// 容易犯的错误：j <= n-1-i。
		// j 表示一个正方形边框中，每一组需要旋转的方格的第一个。如果 j=n-1-i，
		// 那么对于 4x4 的矩阵，i=0 时，j 的取值为 [0, 3]，但是方格 (0,3) 跟
		// 方格 (0,0) 属于一组，这样该组方格就被旋转了两次，结果就不正确了。
		for j := i; j < n-1-i; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = temp
		}
	}
}

func rotate2(matrix [][]int) {
	if matrix == nil {
		return
	}

	// 另一种建模旋转的思路。把矩阵划分成形状相同的 4 部分。旋转第一部分中的每一个单元格，就可以旋转整个矩阵。
	// 观察一个单元格的旋转，会发现这引起 4 个单元格（包含它自身）移动位置。每个单元格移动到下一部分中的对应位置。具体请看
	// https://leetcode-cn.com/problems/rotate-image/solution/xuan-zhuan-tu-xiang-by-leetcode-solution-vu3m/
	// 方法二。
	// n 为偶数时，第一部分包含的单元格数量 n*n/4   = (n/2)   * (n/2)   = (n/2) * (n+1)/2
	// n 为奇数时，第一部分包含的单元格数量 n*n-1/4 = (n-1)/2 * (n+1)/2 = (n/2) * (n+1)/2
	// 两个乘法表达式中的乘数分别表示第一部分的行和列，即外层、内层循环。
	// n 为奇数时，n/2 跟 (n-1)/2 相等。n 为偶数时，n/2 跟 (n+1)/2 相等。
	n := len(matrix)
	for i := 0; i < n/2; i++ { // 第一部分包含的行
		for j := 0; j < (n+1)/2; j++ { // 第一部分包含的列。
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = temp
		}
	}
}

func rotate_flip(matrix [][]int) {
	// TODO 通过翻转实现旋转
	return
}

func rotate_copy(matrix [][]int) {
	n := len(matrix)
	tmp := make([][]int, n)
	for i := range tmp {
		//for i, _ := range tmp {
		tmp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			tmp[j][n-1-i] = matrix[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = tmp[i][j]
		}
	}
}
