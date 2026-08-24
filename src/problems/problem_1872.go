// 1872. Stone Game VIII

package problems

import "fmt"

func Problem_1872() {
	stones := []int{}
	fmt.Println(stoneGameVIII(stones))
}

func stoneGameVIII(stones []int) int {
	for i := 1; i < len(stones); i++ {
		stones[i] += stones[i-1]
	}

	best := stones[len(stones)-1]

	for i := len(stones) - 2; i >= 1; i-- {
		best = max(best, stones[i]-best)
	}

	return best
}
