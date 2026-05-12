// 1665. Minimum Initial Energy to Finish Tasks

package problems

import (
	"fmt"
	"sort"
)

func Problem_1665() {
	tasks := [][]int{
		{1, 2},
		{2, 4},
		{4, 8},
	}
	fmt.Println(minimumEffort(tasks))
}

func minimumEffort(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool {
		return (tasks[i][1] - tasks[i][0]) >= (tasks[j][1] - tasks[j][0])
	})
	fmt.Println(tasks)
	minEnergy := tasks[0][1]
	remEnergy := minEnergy - tasks[0][0]
	for i := 1; i < len(tasks); i++ {
		expendedEnergy := tasks[i][0]
		requiredEnergy := tasks[i][1]
		if remEnergy >= requiredEnergy {
			remEnergy -= expendedEnergy
		} else {
			minEnergy += requiredEnergy - remEnergy
			remEnergy = requiredEnergy - expendedEnergy
		}
	}
	return minEnergy
}
