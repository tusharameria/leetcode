// 1140. Stone Game II

package problems

import "fmt"

func Problem_1140() {
	piles := []int{4, 6, 8, 3, 9, 5, 6, 8, 3, 2, 6}
	fmt.Println(stoneGameII(piles))
}

func stoneGameII(piles []int) int {
	n := len(piles)
	suffix := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suffix[i] = piles[i] + suffix[i+1]
	}
	fmt.Println(suffix)

	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, n+1)
		for j := range n + 1 {
			dp[i][j] = -1
		}
	}

	var solve func(i, m int) int
	solve = func(i, m int) int {
		if i == n {
			return 0
		}
		if dp[i][m] != -1 {
			return dp[i][m]
		}

		best := 0

		for x := 1; x <= 2*m && i+x <= n; x++ {
			nextM := m
			if x > nextM {
				nextM = x
			}

			current := suffix[i] - solve(i+x, nextM)
			if current > best {
				best = current
			}
		}
		dp[i][m] = best
		return best
	}

	return solve(0, 1)
}
