// 1665. Minimum Initial Energy to Finish Tasks

package problems

import (
	"fmt"
)

func Problem_1665() {
	tasks := [][]int{
		{1, 1},
		{1, 3},
	}
	fmt.Println(minimumEffort(tasks))
}

func minimumEffort(tasks [][]int) int {
	diffArr := [10000]int{}
	for _, task := range tasks {
		actualEnergy := task[0]
		minEnergy := task[1]
		diff := minEnergy - actualEnergy
		diffArr[diff] += actualEnergy
	}

	minEnergy := 0
	overflow := 9999
	for {
		if diffArr[overflow] > 0 {
			minEnergy = diffArr[overflow]
			break
		}
		overflow--
	}

	for i := overflow - 1; i >= 0; i-- {
		currentEnergy := diffArr[i]
		minEnergy += currentEnergy
		if currentEnergy+i >= overflow {
			overflow = i
		} else {
			overflow -= currentEnergy
		}
	}

	return minEnergy + overflow
}
