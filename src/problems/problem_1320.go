// 1320. Minimum Distance to Type a Word Using Two Fingers

package problems

import (
	"fmt"
	"slices"
)

func Problem_1320() {
	word := "HAPPY"
	fmt.Println(minimumDistance(word))
}

func minimumDistance(word string) int {
	n := len(word)

	dp := make([]int, 26)

	for i := 1; i < n; i++ {
		prev := int(word[i-1] - 'A')
		curr := int(word[i] - 'A')
		next := make([]int, 26)
		for j := range next {
			next[j] = 1 << 30
		}

		move := dist(prev, curr)
		for other := 0; other < 26; other++ {
			// Move finger at prev -> curr
			next[other] = min(
				next[other],
				dp[other]+move,
			)
			// Move other finger -> curr
			// finger at prev becomes free finger
			next[prev] = min(
				next[prev],
				dp[other]+dist(other, curr),
			)
		}
		dp = next
	}

	return slices.Min(dp)
}

func dist(a, b int) int {
	r1, c1 := a/6, a%6
	r2, c2 := b/6, b%6

	return abs(r1-r2) + abs(c1-c2)
}
