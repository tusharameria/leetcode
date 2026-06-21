// 1833. Maximum Ice Cream Bars

package problems

import (
	"fmt"
)

func Problem_1833() {
	costs := []int{1, 3, 2, 4, 1}
	coins := 20
	fmt.Println(maxIceCream(costs, coins))
}

// Constraints:

// costs.length == n
// 1 <= n <= 10^5
// 1 <= costs[i] <= 10^5
// 1 <= coins <= 10^8

func maxIceCream(costs []int, coins int) int {
	if len(costs) == 0 {
		return 0
	}
	maxCost := 0
	minCost := 1_00_001
	for _, cost := range costs {
		maxCost = max(maxCost, cost)
		minCost = min(minCost, cost)
	}
	store := make([]int, maxCost-minCost+1)
	for _, cost := range costs {
		store[cost-minCost]++
	}
	count := 0
	for i, freq := range store {
		cost := i + minCost
		if coins < cost {
			break
		}
		if freq != 0 {
			for freq > 0 {
				if coins < cost {
					break
				}
				coins -= cost
				count++
				freq--
			}
		}
	}
	return count
}
