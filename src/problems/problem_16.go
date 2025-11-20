package problems

import (
	"fmt"
	"sort"
)

func Problem_16() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	fmt.Println(threeSumClosest(nums, target))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	result := nums[0] + nums[1] + nums[n-1]
	diff := result - target
	if diff < 0 {
		diff = -diff
	}

	for i := 0; i <= n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		current := nums[i]
		left := i + 1
		right := n - 1
		for left < right {
			currentSum := current + nums[left] + nums[right]
			newDiff := currentSum - target
			if newDiff < 0 {
				newDiff = -newDiff
			}

			if newDiff == 0 {
				return target
			} else if newDiff < diff {
				result, diff = currentSum, newDiff
			}

			if currentSum < target {
				left++
			} else {
				right--
			}
		}
	}
	return result
}
