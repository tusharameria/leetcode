package problems

import (
	"fmt"
)

func Problem_knapsack_0_1() {
	weights := []int{2, 1, 3, 2, 4, 5}
	values := []int{12, 10, 20, 15, 25, 30}
	capacity := 7
	fmt.Printf("res : %v\n", knap01(weights, values, capacity, 0))
}

func knap01(weights, values []int, capacity, i int) int {
	if i == len(weights) || capacity <= 0 {
		return 0
	}

	skip := knap01(weights, values, capacity, i+1)
	take := 0

	if weights[i] <= capacity {
		take = values[i] + knap01(weights, values, capacity-weights[i], i+1)
	}

	if take > skip {
		return take
	}
	return skip
}
