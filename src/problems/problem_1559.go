// 1559. Detect Cycles in 2D Grid

package problems

import "fmt"

func Problem_1559() {
	grid := [][]byte{{'a', 'a', 'a', 'a'}, {'a', 'b', 'b', 'a'}, {'a', 'b', 'b', 'a'}, {'a', 'a', 'a', 'a'}}
	fmt.Println(containsCycle(grid))
}

func containsCycle(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var dfs func(y, x, py, px int) bool
	dfs = func(y, x, py, px int) bool {
		visited[y][x] = true

		for _, d := range dirs {
			ny, nx := y+d[0], x+d[1]

			if ny < 0 || nx < 0 || ny >= m || nx >= n {
				continue
			}
			if grid[ny][nx] != grid[y][x] {
				continue
			}
			if ny == py && nx == px {
				continue
			}

			if visited[ny][nx] {
				return true
			}

			if dfs(ny, nx, y, x) {
				return true
			}
		}
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] {
				if dfs(i, j, -1, -1) {
					return true
				}
			}
		}
	}
	return false
}
