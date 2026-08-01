// 486. Predict the Winner

package problems

import (
	"fmt"

	"github.com/tusharameria/leetcode/src/utils"
)

func Problem_486() {
	nums := utils.RandomIntArrayGenerator(1, 23, 12)
	fmt.Println(predictTheWinner(nums))
}

func predictTheWinner(nums []int) bool {
	n := len(nums)
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}

	for k := n - 1; k > 0; k-- {
		for i := 0; i < k; i++ {
			j := i + n - k
			dp[i][j] = max((nums[i] - dp[i+1][j]), (nums[j] - dp[i][j-1]))
		}
	}

	return dp[0][n-1] >= 0
}
