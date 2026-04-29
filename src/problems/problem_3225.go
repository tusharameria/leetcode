// 3225. Maximum Score From Grid Operations

package problems

import "fmt"

func Problem_3225() {
	grid := [][]int{{1, 2}, {3, 4}}
	fmt.Println(maximumScore(grid))
}

func maximumScore(grid [][]int) int64 {
	n := len(grid)
	prefixCol := make([][]int64, n+1)
	for r := range prefixCol {
		prefixCol[r] = make([]int64, n)
	}
	for c := 0; c < n; c++ {
		for r := 0; r < n; r++ {
			prefixCol[r+1][c] = prefixCol[r][c] + int64(grid[r][c])
		}
	}

	dp := make([][][]int64, n+1)
	for i := range dp {
		dp[i] = make([][]int64, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int64, n+1)
		}
	}

	for colIdx := n; colIdx >= 0; colIdx-- {
		// Pre-calculating the best row choice for every p1 to remove the inner i-loop
		bestPre := make([][]int64, n+1)
		bestSuf := make([][]int64, n+1)
		if colIdx < n {
			for p1 := -1; p1 < n; p1++ {
				bestPre[p1+1] = make([]int64, n+1)
				bestSuf[p1+1] = make([]int64, n+1)

				var curPre int64 = 0
				for i := 0; i < n; i++ {
					curPre = max(curPre, dp[colIdx+1][i+1][p1+1])
					bestPre[p1+1][i] = curPre
				}

				var curSuf int64 = -1e18
				for i := n - 1; i >= 0; i-- {
					val := dp[colIdx+1][i+1][p1+1]
					if colIdx > 0 {
						val += prefixCol[i+1][colIdx-1]
					}
					curSuf = max(curSuf, val)
					bestSuf[p1+1][i] = curSuf
				}
			}
		}

		for p1 := -1; p1 < n; p1++ {
			for p2 := -1; p2 < n; p2++ {
				// Simplified initial sum using prefixCol
				sum := int64(0)
				if p2 > p1 && colIdx > 0 {
					sum = prefixCol[p2+1][colIdx-1] - prefixCol[p1+1][colIdx-1]
				}

				if colIdx == n {
					dp[colIdx][p1+1][p2+1] = sum
					continue
				}

				// res for no coloring
				res := sum + dp[colIdx+1][0][p1+1]

				// Optimized lookup for coloring rows up to i
				highestPrev := max(p1, p2)
				var bestScore int64 = 0
				if highestPrev >= 0 {
					bestScore = max(bestScore, bestPre[p1+1][highestPrev])
				}
				if highestPrev < n-1 {
					val := bestSuf[p1+1][highestPrev+1]
					if colIdx > 0 {
						val -= prefixCol[highestPrev+1][colIdx-1]
					}
					bestScore = max(bestScore, val)
				} else if highestPrev == -1 {
					bestScore = max(bestScore, bestSuf[p1+1][0])
				}

				dp[colIdx][p1+1][p2+1] = max(res, sum+bestScore)
			}
		}
	}
	return dp[0][0][0]
}
