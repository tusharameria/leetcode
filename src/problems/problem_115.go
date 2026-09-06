// 115.Distinct Subsequences

package problems

import "fmt"

func Problem_115() {
	s := "hjhjds"
	t := "hjd"
	fmt.Println(numDistinct(s, t))
}

func numDistinct(s, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}
	dp := make([][]int32, m+1)
	for i := range dp {
		dp[i] = make([]int32, n+1)
		dp[i][n] = 1
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s[i] == t[j] {
				dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	return int(dp[0][0])
}
