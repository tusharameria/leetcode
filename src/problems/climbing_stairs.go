package problems

import "fmt"

func ClimbingStairs() {
	// find number of ways in which someone can climb the stairs if they can take either 1 or 2 steps
	n := 60
	fmt.Println(countWays(n))
}

func countWays(n int) int {
	if n <= 2 {
		return n
	}

	prev1, prev2 := 1, 2
	for i := 3; i <= n; i++ {
		prev1, prev2 = prev2, prev1+prev2
	}
	return prev2
}
