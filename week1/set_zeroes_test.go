package week1

import (
	"fmt"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	//matrix := [][]int{
	//	{0, 3, -1},
	//	{2, 5, -2},
	//	{1, 4, 0},
	//}
	//setZeroes(matrix)
	//fmt.Println(matrix)

	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix)
	fmt.Println(matrix)
}

func TestSetZeroV2(t *testing.T) {
	matrix := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	setZeroesV2_bug(matrix)
	fmt.Println(matrix)
}
