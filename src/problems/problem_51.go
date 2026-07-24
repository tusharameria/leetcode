// 51. N-Queens

package problems

import "fmt"

func Problem_51() {
	n := 4
	fmt.Println(solveNQueens(n))
}

func solveNQueens(n int) [][]string {
	grid := make([][]int8, n)
	for i := range n {
		grid[i] = make([]int8, n)
	}
	return nil
}
