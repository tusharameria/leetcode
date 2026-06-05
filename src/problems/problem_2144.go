// 2144. Minimum Cost of Buying Candies With Discount

package problems

import (
	"fmt"
	"sort"
)

func Problem_2144() {
	cost := []int{6, 5, 7, 9, 2, 2}
	fmt.Println(minimumCost(cost))
}

func minimumCost(cost []int) int {
	sort.Ints(cost)
	n := len(cost)
	sum := 0
	if n <= 2 {
		for _, val := range cost {
			sum += val
		}
	} else {
		rem := (n - 3) % 3
		for i := n - 1; i >= 0; i-- {
			if i%3 != rem {
				sum += cost[i]
			}
		}
	}
	return sum
}
