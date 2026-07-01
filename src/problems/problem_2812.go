// 2812. Find the Safest Path in a Grid

package problems

import (
	"fmt"
)

func Problem_2812() {
	grid := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0},
	}
	fmt.Println(maximumSafenessFactor(grid))
}

func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)
	manGrid := make([][]int, n)
	onesCount := 0
	stack := make([][2]int, n*n)
	for i := range manGrid {
		manGrid[i] = make([]int, n)
		for j := range manGrid[i] {
			manGrid[i][j] = -1
			if grid[i][j] == 1 {
				stack[onesCount] = [2]int{i, j}
				onesCount++
				manGrid[i][j] = 0
			}
		}
	}

	// for _, row := range manGrid {
	// 	fmt.Println(row)
	// }

	stack = stack[:onesCount]

	maxDist := 0
	for len(stack) > 0 {
		curr := stack[0]
		stack = stack[1:]
		x, y := curr[0], curr[1]
		if maxDist < manGrid[x][y] {
			maxDist = manGrid[x][y]
		}
		left := [2]int{x - 1, y}
		up := [2]int{x, y - 1}
		right := [2]int{x + 1, y}
		down := [2]int{x, y + 1}

		if left[0] >= 0 && manGrid[left[0]][left[1]] == -1 {
			manGrid[left[0]][left[1]] = manGrid[x][y] + 1
			stack = append(stack, left)
		}

		if up[1] >= 0 && manGrid[up[0]][up[1]] == -1 {
			manGrid[up[0]][up[1]] = manGrid[x][y] + 1
			stack = append(stack, up)
		}

		if right[0] < n && manGrid[right[0]][right[1]] == -1 {
			manGrid[right[0]][right[1]] = manGrid[x][y] + 1
			stack = append(stack, right)
		}

		if down[1] < n && manGrid[down[0]][down[1]] == -1 {
			manGrid[down[0]][down[1]] = manGrid[x][y] + 1
			stack = append(stack, down)
		}
	}

	// fmt.Println("max : ", maxDist)

	// for _, row := range manGrid {
	// 	fmt.Println(row)
	// }

	s, e := 0, min(maxDist, manGrid[0][0], manGrid[n-1][n-1])
	for s <= e {
		// fmt.Printf("-- s : %d, e : %d ----", s, e)
		mid := (s + e) / 2
		isPossible := false
		visited := make([][]bool, n)
		for i := 0; i < n; i++ {
			visited[i] = make([]bool, n)
		}
		stack := [][2]int{{0, 0}}
		visited[0][0] = true
		for len(stack) > 0 {
			curr := stack[0]
			stack = stack[1:]
			x, y := curr[0], curr[1]
			if x == n-1 && y == n-1 {
				isPossible = true
				break
			}

			left := [2]int{x - 1, y}
			up := [2]int{x, y - 1}
			right := [2]int{x + 1, y}
			down := [2]int{x, y + 1}

			if left[0] >= 0 && !visited[left[0]][left[1]] && manGrid[left[0]][left[1]] >= mid {
				stack = append(stack, left)
				visited[left[0]][left[1]] = true
			}
			if up[1] >= 0 && !visited[up[0]][up[1]] && manGrid[up[0]][up[1]] >= mid {
				stack = append(stack, up)
				visited[up[0]][up[1]] = true
			}
			if right[0] < n && !visited[right[0]][right[1]] && manGrid[right[0]][right[1]] >= mid {
				stack = append(stack, right)
				visited[right[0]][right[1]] = true
			}
			if down[1] < n && !visited[down[0]][down[1]] && manGrid[down[0]][down[1]] >= mid {
				stack = append(stack, down)
				visited[down[0]][down[1]] = true
			}
			// fmt.Printf("-- s : %d, e : %d ----", s, e)
			// fmt.Println(stack)
			// time.Sleep(time.Millisecond * 100)
		}

		if isPossible {
			s = mid + 1
		} else {
			e = mid - 1
		}
	}

	return e
}
