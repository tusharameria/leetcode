// 396. Rotate Function

package problems

import "fmt"

func Problem_396() {
	nums := []int{4, 3, 2, 6}
	fmt.Println("res : ", maxRotateFunction(nums))
}

func maxRotateFunction(nums []int) int {
	numLen := len(nums)
	if numLen <= 1 {
		return 0
	}
	wholeSum := 0
	max := 0
	currentSum := 0
	for i := 0; i < numLen; i++ {
		wholeSum += nums[i]
		currentSum += i * nums[i]
	}
	max = currentSum
	for i := numLen - 1; i >= 1; i-- {
		currentSum += wholeSum - (numLen * nums[i])
		if currentSum > max {
			max = currentSum
		}
	}
	return max
}
