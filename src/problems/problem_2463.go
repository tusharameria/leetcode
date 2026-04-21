// 2463. Minimum Total Distance Traveled

package problems

import (
	"fmt"
	"sort"
)

func Problem_2463() {
	robot := []int{1, -1}
	factory := [][]int{
		{-2, 1},
		{2, 1},
	}
	fmt.Println(minimumTotalDistance(robot, factory))
}

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	sort.Ints(robot)
	sort.Slice(factory, func(i, j int) bool {
		return factory[i][0] < factory[j][0]
	})

	n, m := len(robot), len(factory)
	const INF int64 = 1e18

	dp := make([][]int64, n+1)
	for i := range dp {
		dp[i] = make([]int64, m+1)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}

	// base case
	for j := 0; j <= m; j++ {
		dp[0][j] = 0
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// skip this factory
			dp[i][j] = dp[i][j-1]

			cost := int64(0)

			for k := 1; k <= factory[j-1][1] && i-k >= 0; k++ {
				pos := factory[j-1][0]
				cost += int64(abs(robot[i-k] - pos))

				if dp[i-k][j-1]+cost < dp[i][j] {
					dp[i][j] = dp[i-k][j-1] + cost
				}
			}
		}
	}

	return dp[n][m]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
