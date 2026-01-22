package problems

import (
	"fmt"
)

// 3507. Minimum Pair Removal to Sort Array I

func Problem_3507() {
	nums := []int{1, 2, 2}
	fmt.Printf("minimumPairRemoval : %v\n", minimumPairRemoval(nums))
}

func minimumPairRemoval(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	count := 0
	min := nums[1] + nums[0]
	minIndex := 1
	nonDec := true
	for i := 1; i <= len(nums)-1; i++ {
		if nums[i] < nums[i-1] {
			nonDec = false
		}
		if nums[i]+nums[i-1] < min {
			min = nums[i] + nums[i-1]
			minIndex = i
		}
		if i == len(nums)-1 && !nonDec {
			count++
			nums = append(nums[:minIndex-1], append([]int{min}, nums[minIndex+1:]...)...)
			i = 0
			if len(nums) <= 1 {
				break
			}
			min = nums[1] + nums[0]
			minIndex = 1
			nonDec = true
		}
	}
	return count
}
