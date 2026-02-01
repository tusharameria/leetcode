package problems

import "fmt"

// 3010. Divide an Array Into Subarrays With Minimum Cost I

func Problem_3010() {
	nums := []int{10, 3, 1, 1}
	fmt.Printf("minimumCost : %v\n", minimumCost(nums))
}

func minimumCost(nums []int) int {
	lens := len(nums)
	secondMin := nums[1]
	min := nums[2]
	if secondMin < min {
		secondMin, min = min, secondMin
	}

	for i := lens - 1; i >= 3; i-- {
		if nums[i] <= min {
			min, secondMin = nums[i], min
		} else if nums[i] < secondMin {
			secondMin = nums[i]
		}
	}

	return nums[0] + secondMin + min
}
