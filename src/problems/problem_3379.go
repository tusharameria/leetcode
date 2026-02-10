package problems

import (
	"fmt"
	"math"
)

func Problem_3379() {
	nums := []int{3, -2, 1, 1}
	fmt.Printf("res : %v\n", constructTransformedArray(nums))
}

func constructTransformedArray(nums []int) []int {
	lens := len(nums)
	res := make([]int, lens)
	for i := 0; i <= lens-1; i++ {
		if nums[i] == 0 {
			res[i] = nums[i]
		} else if nums[i] > 0 {
			res[i] = nums[(i+nums[i])%lens]
		} else {
			index := (i - int(math.Abs(float64(nums[i])))) % lens
			if index < 0 {
				index += lens
			}
			res[i] = nums[index]
		}
	}
	return res
}
