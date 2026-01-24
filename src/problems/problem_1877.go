package problems

import (
	"fmt"
	"sort"
)

// 1877. Minimize Maximum Pair Sum in Array

func Problem_1877() {
	nums := []int{3, 5, 4, 2, 4, 6}
	fmt.Printf("minPairSum : %d\n", minPairSum(nums))
}

func minPairSum(nums []int) int {
	sort.Ints(nums)
	if len(nums) == 2 {
		return nums[0] + nums[1]
	}
	maxSum := 0
	for i := 0; i <= (len(nums)/2)-1; i++ {
		maxSum = max(maxSum, nums[i]+nums[len(nums)-i-1])
	}
	return maxSum
}
