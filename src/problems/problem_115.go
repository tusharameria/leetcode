// 115.Distinct Subsequences

package problems

import "fmt"

func Problem_115() {
	s := "eee"
	t := "eee"
	fmt.Println(numDistinct(s, t))
}

func numDistinct(s, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}
	dp := make([]int32, m+1)
	for i := m - 1; i >= 0; i-- {
		dp[i] = dp[i+1]
		if s[i] == t[n-1] {
			dp[i]++
		}
	}
	fmt.Println(dp)
	dp[m] = 0
	for j := n - 2; j >= 0; j-- {
		prevVal := int32(0)
		for i := m - 1; i >= 0; i-- {
			buff := dp[i]
			dp[i] = dp[i+1]
			if s[i] == t[j] {
				dp[i] += prevVal
			}
			prevVal = buff
		}
		fmt.Println(dp)
	}
	return int(dp[0])
}
