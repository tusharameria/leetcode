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
	y := len(grid)
	x := len(grid[0])

	res := make([][]int, y)
	for i := range y {
		res[i] = make([]int, x)
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			ny := (k / x) + i
			if (j + (k % x)) >= x {
				ny++
			}
			res[ny%y][(j+(k%x))%x] = grid[i][j]
		}
	}

	return res
}
