// 877. Stone Game

package problems

import (
	"fmt"

	"github.com/tusharameria/leetcode/src/utils"
)

func Problem_877() {
	piles := utils.RandomIntArrayGenerator(1, 23, 12)
	fmt.Println(stoneGame(piles))
}

func stoneGame(piles []int) bool {
	n := len(piles)
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, n)
		dp[i][i] = piles[i]
	}

	for k := n - 1; k > 0; k-- {
		for i := 0; i < k; i++ {
			j := i + n - k
			dp[i][j] = max((piles[i] - dp[i+1][j]), (piles[j] - dp[i][j-1]))
		}
	}

	return dp[0][n-1] >= 0
}
