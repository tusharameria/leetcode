// 41. First Missing Positive

package problems

import "fmt"

func Problem_41() {
	nums := []int{100000, 3, 4000, 2, 15, 1, 99999}
	fmt.Println(firstMissingPositive(nums))
}

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		targetIndex := nums[i] - 1
		for targetIndex >= 0 && targetIndex < n {
			targetIndexVal := nums[targetIndex]
			if targetIndexVal == targetIndex+1 {
				break
			}
			nums[i], nums[targetIndex] = targetIndexVal, targetIndex+1
			targetIndex = targetIndexVal - 1
		}
	}

	for i, val := range nums {
		if i+1 != val {
			return i + 1
		}
	}
	return n + 1
}
