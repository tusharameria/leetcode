// 1140. Stone Game II

package problems

import (
	"fmt"

	"github.com/tusharameria/leetcode/src/utils"
)

func Problem_1140() {
	for i := 0; i < 1; i++ {
		piles := utils.RandomIntArrayGenerator(1, 100, 100)
		fmt.Println(stoneGameII(piles))
	}
}

const STEP = 1 << 20
const MASK = STEP - 1

var dp [101][101]int
var suf [101]int
var gen int

func solve(i, m, n int) (win int) {
	if i >= n {
		return 0
	} else if 2*m >= n-i {
		return suf[i]
	} else if dp[i][m] >= gen {
		return dp[i][m] & MASK
	}
	for j := 1; j <= 2*m; j++ {
		win = max(win, suf[i]-solve(i+j, max(m, j), n))
	}
	dp[i][m] = win | gen
	return win
}

func stoneGameII(piles []int) int {
	n := len(piles)
	gen += STEP
	suf[n] = 0
	for i := n - 1; i >= 0; i-- {
		suf[i] = suf[i+1] + piles[i]
	}
	return solve(0, 1, n)
}
