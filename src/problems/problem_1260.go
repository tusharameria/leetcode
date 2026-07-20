// 1260. Shift 2D Grid

package problems

import "fmt"

func Problem_1260() {
	grid := [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	k := 1
	for _, row := range shiftGrid(grid, k) {
		fmt.Println(row)
	}
}

func shiftGrid(grid [][]int, k int) [][]int {
	rowLen := len(grid)
	colLen := len(grid[0])
	k = k % (rowLen * colLen)

	if k == 0 {
		return grid
	}

	res := make([][]int, rowLen)
	for i := range rowLen {
		res[i] = make([]int, colLen)
	}

	idx := 0
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			newIdx := (idx + k) % (rowLen * colLen)
			res[newIdx/colLen][newIdx%colLen] = grid[i][j]
			idx++
		}
	}

	return res
}
