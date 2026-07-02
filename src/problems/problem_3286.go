// 3286. Find a Safe Walk Through a Grid

package problems

import (
	"fmt"
)

func Problem_3286() {
	grid := [][]int{
		{0, 1, 1, 0, 0, 0},
		{1, 0, 1, 0, 0, 0},
		{0, 1, 1, 1, 0, 1},
		{0, 0, 1, 0, 1, 0},
	}
	health := 3
	fmt.Println(findSafeWalk(grid, health))
}

func findSafeWalk(grid [][]int, health int) bool {
	rowLen := len(grid)
	colLen := len(grid[0])
	healthGrid := make([][]int, rowLen)
	for i := range rowLen {
		healthGrid[i] = make([]int, colLen)
		for j := range colLen {
			healthGrid[i][j] = -1
		}
	}

	directions := [4][2]int{
		{-1, 0},
		{0, -1},
		{1, 0},
		{0, 1},
	}

	stack := [][2]int{{0, 0}}
	healthGrid[0][0] = health

	for len(stack) > 0 {
		curr := stack[0]
		stack = stack[1:]

		for _, dirVal := range directions {
			x, y := curr[0]+dirVal[0], curr[1]+dirVal[1]
			if x >= 0 && x < rowLen && y >= 0 && y < colLen {
				// fmt.Printf("- checking [%d, %d] -\n", x, y)
				targetHealth := healthGrid[curr[0]][curr[1]]
				if grid[x][y] == 1 {
					targetHealth--
				}
				if targetHealth > 0 && targetHealth > healthGrid[x][y] {
					stack = append(stack, [2]int{x, y})
					healthGrid[x][y] = targetHealth
				}
			}
			// time.Sleep(time.Millisecond * 500)
		}
	}
	for _, row := range healthGrid {
		fmt.Println(row)
	}

	return healthGrid[rowLen-1][colLen-1] >= 1
}
