// 3742. Maximum Path Score in a Grid

package problems

import "fmt"

func Problem_3742() {
	grid := [][]int{{1, 2}, {0, 2}}
	fmt.Println(maximumPathScore(grid, 2))
}

func maximumPathScore(grid [][]int, k int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][][]int, m+1)
	INFINT := 1 << 30
	for i := 0; i < m+1; i++ {
		dp[i] = make([][]int, n+1)
		for j := 0; j < n+1; j++ {
			dp[i][j] = make([]int, k+1)
			for l := 0; l < k+1; l++ {
				dp[i][j][l] = -INFINT
			}
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for l := 0; l <= k; l++ {
				if i == 1 && j == 1 {
					dp[i][j][l] = 0
				} else {
					score := grid[i-1][j-1]
					cost := (score + 1) / 2
					if l-cost < 0 {
						dp[i][j][l] = -INFINT
						continue
					}
					upperScore := dp[i-1][j][l-cost]
					leftScore := dp[i][j-1][l-cost]
					maxScore := max(upperScore, leftScore)
					if maxScore == -INFINT {
						dp[i][j][l] = -INFINT
						continue
					}
					dp[i][j][l] = maxScore + score
				}
			}
		}
	}
	finalMaxScore := -INFINT
	for i := k; i >= 0; i-- {
		curr := dp[m][n][i]
		if curr == -INFINT {
			break
		}
		finalMaxScore = max(curr, finalMaxScore)
	}
	if finalMaxScore == -INFINT {
		return -1
	}
	return finalMaxScore
}
