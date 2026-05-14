// 1674. Minimum Moves to Make Array Complementary

package problems

import (
	"fmt"
)

func Problem_1674() {
	nums := []int{1, 2, 4, 3}
	limit := 4
	fmt.Println(minMoves(nums, limit))
}

func minMoves(nums []int, limit int) int {
	minSum := 0
	maxSum := 0
	for i := 0; i < len(nums)/2; i++ {
		sum := nums[i] + nums[len(nums)-1-i]
		minSum = min(minSum, sum)
		maxSum = max(maxSum, sum)
	}
	if minSum == maxSum {
		return 0
	}
	counts := make([]int, maxSum-minSum+1)

	for i := 0; i < len(nums)/2; i++ {
		minNum := min(nums[i], nums[len(nums)-1-i])
		maxNum := max(nums[i], nums[len(nums)-1-i])
		sum := nums[i] + nums[len(nums)-1-i]
		// min with 1 change
		min1Change := minNum + 1
		max1Change := maxNum + limit
		if minNum >= minSum {
			counts[0] += 2
			counts[min1Change-minSum] -= 2
		}
		counts[max(0, min1Change-minSum)] += 1
		if max1Change+1 <= maxSum {
			counts[max1Change+1-minSum] += 1
		}
		counts[sum-minSum] -= 1
		if sum != maxSum {
			counts[sum+1-minSum] += 1
		}
	}
	res := counts[0]
	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
		res = min(res, counts[i])
	}
	return res
}
