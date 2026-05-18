// 154. Find Minimum in Rotated Sorted Array II

package problems

import "fmt"

func Problem_154() {
	// nums := []int{1, 2, 1}
	// nums := []int{2, 1, 2}
	nums := []int{10, 1, 10, 10, 10}
	fmt.Println(findMin(nums))
}

func findMin(nums []int) int {
	n := len(nums)
	left := 0
	right := n - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] {
			right = mid
		} else if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right--
		}
	}
	return nums[right]
}
