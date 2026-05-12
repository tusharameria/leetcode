// 3418. Maximum Amount of Money Robot Can Earn

package problems

import "fmt"

func Problem_3418() {
	events := [][]int{
		{0, 1, -1},
		{1, -2, 3},
		{2, -3, 4},
	}
	fmt.Println(maximumAmount(events))
}

func maximumAmount(coins [][]int) int {
	height := len(coins)
	width := len(coins[0])
	dp := make([][]int, width+1)
	for i := range dp {
		dp[i] = make([]int, 3)
	}
	for j := 0; j <= width-1; j++ {
		val := coins[0][j]
		if val >= 0 {
			for k := 0; k < 3; k++ {
				dp[j+1][k] = dp[j][k] + val
			}
		} else {
			for k := 2; k > 0; k-- {
				dp[j+1][k] = max(dp[j][k]+val, dp[j][k-1])
			}
			dp[j+1][0] = dp[j][0] + val
		}
	}

	for i := 1; i <= height-1; i++ {
		val0 := coins[i][0]
		if val0 >= 0 {
			for k := 0; k < 3; k++ {
				dp[1][k] += val0
			}
		} else {
			for k := 2; k > 0; k-- {
				dp[1][k] = max(dp[1][k]+val0, dp[1][k-1])
			}
			dp[1][0] += val0
		}
		for j := 1; j <= width-1; j++ {
			val := coins[i][j]
			if val >= 0 {
				for k := 0; k < 3; k++ {
					dp[j+1][k] = max(dp[j+1][k], dp[j][k]) + val
				}
			} else {
				for k := 2; k > 0; k-- {
					dp[j+1][k] = max(dp[j+1][k]+val, dp[j+1][k-1], dp[j][k]+val, dp[j][k-1])
				}
				dp[j+1][0] = max(dp[j+1][0], dp[j][0]) + val
			}
		}
	}
	return dp[width][2]
}
