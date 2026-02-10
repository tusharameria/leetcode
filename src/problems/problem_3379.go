package problems

import (
	"fmt"
)

func Problem_3379() {
	nums := []int{3, -2, 1, 1}
	fmt.Printf("res : %v\n", constructTransformedArray(nums))
}

func constructTransformedArray(nums []int) []int {
	lens := len(nums)
	res := make([]int, lens)
	for i := 0; i <= lens-1; i++ {
		index := i
		if nums[i] > 0 {
			index = (i + nums[i]) % lens
		} else if nums[i] < 0 {
			index = (i + nums[i]) % lens
			if index < 0 {
				index += lens
			}
		}
		res[i] = nums[index]
	}
	return res
}
