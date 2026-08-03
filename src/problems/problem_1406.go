// 1406. Stone Game III

package problems

import (
	"fmt"
)

func Problem_1406() {
	// stoneValue := utils.RandomIntArrayGenerator(-30, 30, 10)
	stoneValue := []int{-10, -9, -6, -11, 12, -15, 5, 4, -9, 9, 6, -5, 12, 16, -3, 9, 0, 13, 1, -10, 6, -14, 13}
	fmt.Println(stoneGameIII(stoneValue))
}

const ALICE = "Alice"
const BOB = "Bob"
const TIE = "Tie"

func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)
	if n == 1 {
		if stoneValue[0] > 0 {
			return ALICE
		} else if stoneValue[0] < 0 {
			return BOB
		} else {
			return TIE
		}
	}
	dp := make([][]int, 3)
	for i := range 3 {
		dp[i] = make([]int, n)
	}
	dp[0][n-1] = stoneValue[n-1]
	for i := n - 2; i >= 0; i-- {
		dp[0][i] = stoneValue[i] - dp[0][i+1]
	}
	dp[1][n-1] = dp[0][n-1]
	dp[1][n-2] = max(
		stoneValue[n-2]+stoneValue[n-1],
		stoneValue[n-2]-dp[1][n-1],
	)
	if n == 2 {
		if dp[1][n-2] > 0 {
			return ALICE
		} else if dp[1][n-2] < 0 {
			return BOB
		} else {
			return TIE
		}
	}
	for i := n - 3; i >= 0; i-- {
		dp[1][i] = max(
			stoneValue[i]+stoneValue[i+1]-dp[1][i+2],
			stoneValue[i]-dp[1][i+1],
		)
	}
	dp[2][n-1] = dp[1][n-1]
	dp[2][n-2] = dp[1][n-2]
	dp[2][n-3] = max(
		stoneValue[n-3]+stoneValue[n-2]+stoneValue[n-1],
		stoneValue[n-3]+stoneValue[n-2]-dp[2][n-1],
		stoneValue[n-3]-dp[2][n-2],
	)
	for i := n - 4; i >= 0; i-- {
		dp[2][i] = max(
			stoneValue[i]+stoneValue[i+1]+stoneValue[i+2]-dp[2][i+3],
			stoneValue[i]+stoneValue[i+1]-dp[2][i+2],
			stoneValue[i]-dp[2][i+1],
		)
	}
	for _, row := range dp {
		fmt.Println(row)
	}
	if dp[2][0] > 0 {
		return ALICE
	} else if dp[2][0] < 0 {
		return BOB
	}
	return TIE
}
