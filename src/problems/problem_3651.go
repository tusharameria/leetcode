// 3651. Minimum Cost Path with Teleportations

package problems

import (
	"fmt"
	"slices"

	"github.com/tusharameria/leetcode/src/utils"
)

func Problem_3651() {
	grid := utils.RandomInt2DArrayGenerator(1, 20, 3, 4)
	grid = [][]int{{1, 3, 3}, {2, 5, 4}, {4, 3, 5}}
	k := 2
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println(minCost(grid, k))
}

type Cell struct {
	val, r, c int
}

func minCost(grid [][]int, k int) int {
	row := len(grid)
	col := len(grid[0])
	if grid[0][0] >= grid[row-1][col-1] && k >= 1 {
		return 0
	}

	const inf int = 1 << 60

	// ------------------------
	// ------------------------
	// DP with 0 jumps
	// ------------------------
	prev := make([][]int, row)
	for i := 0; i < row; i++ {
		prev[i] = make([]int, col)
		for j := 0; j < col; j++ {
			prev[i][j] = inf
		}
	}

	prev[0][0] = 0

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i > 0 {
				prev[i][j] = min(prev[i][j], prev[i-1][j]+grid[i][j])
			}
			if j > 0 {
				prev[i][j] = min(prev[i][j], prev[i][j-1]+grid[i][j])
			}
		}
	}

	res := prev[row-1][col-1]

	fmt.Printf("----- %d Jumps -----\n", 0)
	for _, row := range prev {
		fmt.Println(row)
	}

	// Sort the complete grid descending with coordinates
	cells := make([]Cell, row*col)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			cells[(i*col)+j] = Cell{
				val: grid[i][j],
				r:   i,
				c:   j,
			}
		}
	}

	slices.SortFunc(cells, func(a, b Cell) int {
		switchVal := b.val > a.val
		switch {
		case switchVal:
			return 1
		default:
			return -1
		}
	})

	fmt.Println("Cells")
	fmt.Println(cells)

	for jump := 1; jump < k+1; jump++ {
		curr := make([][]int, row)
		for i := range row {
			curr[i] = make([]int, col)
			for j := range col {
				curr[i][j] = inf
			}
		}

		// From the cell with highest value, we can go anywhere else in the grid
		best := inf
		// From the cell with 2nd highest value, we can go anywhere else in the grid less or equal to that val
		// And so on

		// So we can update the best value prefix wise
		// Best value would be minimum of
		// 1. current value from k-1 jumps at this coordinate
		// 2. current value from k-1 jumps from previous higher value coordinate

		for i := 0; i < row*col; {
			j := i
			for j < row*col && cells[i].val == cells[j].val {
				r := cells[j].r
				c := cells[j].c
				best = min(best, prev[r][c])
				j++
			}
			for p := i; p < j; p++ {
				r := cells[p].r
				c := cells[p].c
				curr[r][c] = best
			}
			i = j
		}

		fmt.Printf("------------------------------\n")
		fmt.Printf("----- %d Jumps -----\n", jump)
		fmt.Printf("------------------------------\n")
		fmt.Printf("----- After Teleportation -----\n")
		for _, row := range curr {
			fmt.Println(row)
		}

		for i := 0; i < row; i++ {
			for j := 0; j < col; j++ {
				if i > 0 {
					curr[i][j] = min(curr[i][j], curr[i-1][j]+grid[i][j])
				}
				if j > 0 {
					curr[i][j] = min(curr[i][j], curr[i][j-1]+grid[i][j])
				}
			}
		}

		res = min(res, curr[row-1][col-1])
		prev = curr

		fmt.Printf("----- After Right/Down -----\n")
		for _, row := range prev {
			fmt.Println(row)
		}
	}

	return res
}
