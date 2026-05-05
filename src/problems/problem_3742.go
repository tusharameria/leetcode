// 3742. Maximum Path Score in a Grid

package problems

import (
	"fmt"
	"slices"
)

func Problem_3742() {
	grid := [][]int{{0, 1, 2}, {0, 1, 1}, {2, 0, 2}}
	fmt.Println(maxPathScore(grid, 2))
}

func maxPathScore(grid [][]int, k int) int {
	m := len(grid)
	n := len(grid[0])
	var dp [201][400]int
	INFINT := 1 << 30

	// Check if reaching end tile is possible
	firstRow := dp[0]
	for j, val := range grid[0] {
		firstRow[j+1] = firstRow[j] + min(val, 1)
	}
	firstRow[0] = INFINT
	for i := 1; i < m; i++ {
		for j, val := range grid[i] {
			firstRow[j+1] = min(firstRow[j], firstRow[j+1]) + min(val, 1)
		}
	}

	if firstRow[n] > k {
		return -1
	}

	// Max Score
	for j := range n + 1 {
		for c := range min(k+2, m+n) {
			dp[j][c] = -INFINT
		}
	}

	dp[1][1] = 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			val := grid[i][j]
			step := 1 - min(1, val)
			for c := min(k+2, i+j); c >= 0; c-- {
				dp[j+1][c+1] = max(dp[j+1][c+step], dp[j][c+step]) + val
			}
		}
	}

	return slices.Max(dp[n][:min(k+2, m+n)])
}
