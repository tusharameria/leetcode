// 2574. Left and Right Sum Differences

package problems

import "fmt"

func Problem_2574() {
	nums := []int{10, 4, 8, 3}
	fmt.Println(leftRightDifference(nums))
}

func leftRightDifference(nums []int) []int {
	n := len(nums)
	leftSum := 0
	rightSum := 0
	for _, val := range nums {
		rightSum += val
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		val := nums[i]
		rightSum -= val
		ans[i] = abs(leftSum - rightSum)
		leftSum += val
	}
	return ans
}
