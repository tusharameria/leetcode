package problems

import "fmt"

// 3190. Find Minimum Operations to Make All Elements Divisible by Three

func Problem_3190() {
	nums := []int{1, 2, 3, 4}
	fmt.Printf("res : %d\n", minimumOperations(nums))
}

func minimumOperations(nums []int) int {
	ops := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%3 != 0 {
			ops++
		}
	}
	return ops
}
