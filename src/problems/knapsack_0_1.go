package problems

import (
	"fmt"
)

func Problem_knapsack_0_1() {
	weights := []int{2, 1, 3, 2, 4, 5}
	values := []int{12, 10, 20, 15, 25, 30}
	capacity := 7
	fmt.Printf("res : %v\n", knap01(weights, values, capacity))
}

func knap01(weights, values []int, capacity int) int {
	dp := make([][]int, len(weights))
	for i := 0; i <= len(dp)-1; i++ {
		dp[i] = make([]int, capacity+1)
	}
	for i := len(dp) - 1; i >= 0; i-- {
		for j := 1; j <= capacity; j++ {
			inc := 0
			if weights[i] <= j {
				inc = values[i]
			}
			if j-weights[i] >= 0 {
				if i+1 < len(dp) {
					inc += dp[i+1][j-weights[i]]
				}
			}
			exc := 0
			if i+1 < len(dp) {
				exc = dp[i+1][j]
			}
			dp[i][j] = max(inc, exc)
		}
	}
	return dp[0][capacity]
}
