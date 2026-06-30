// 1582. Special Positions in a Binary Matrix

package problems

import (
	"fmt"

	"github.com/tusharameria/leetcode/src/utils"
)

func Problem_1582() {
	grid := utils.RandomInt2DArrayGenerator(0, 1, 5, 5)
	grid = [][]int{
		{1, 0, 0},
		{0, 0, 1},
		{1, 0, 0},
	}
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println("------------")
	fmt.Println(numSpecial(grid))
}

func numSpecial(mat [][]int) int {
	rowLen := len(mat)
	colLen := len(mat[0])
	rowsCount := make([]int, rowLen)
	colsCount := make([]int, colLen)
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if mat[i][j] == 1 {
				rowsCount[i]++
			}
		}
	}
	for j := 0; j < colLen; j++ {
		for i := 0; i < rowLen; i++ {
			if mat[i][j] == 1 {
				colsCount[j]++
			}
		}
	}
	// fmt.Println("rows")
	// fmt.Println(rowsCount)
	// fmt.Println("rowsIdx")
	// fmt.Println(rowsIdx)
	// fmt.Println("cols")
	// fmt.Println(colsCount)
	// fmt.Println("colsIdx")
	// fmt.Println(colsIdx)

	res := 0
	for i := 0; i < rowLen; i++ {
		if rowsCount[i] == 1 {
			for j := 0; j < colLen; j++ {
				if mat[i][j] == 1 && colsCount[j] == 1 {
					res++
					break
				}
			}
		}
	}

	return res
}
