// 153. Find Minimum in Rotated Sorted Array

package problems

import (
	"fmt"
)

func Problem_153() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(findMin(nums))
}

func findMin(nums []int) int {
	n := len(nums)
	left := 0
	right := n - 1
	leftVal := nums[left]
	rightVal := nums[right]
	for left < right {
		mid := left + (right-left)/2
		midVal := nums[mid]
		if leftVal <= rightVal {
			return leftVal
		} else {
			if leftVal <= midVal {
				left = mid + 1
				leftVal = nums[left]
			} else {
				right = mid
				rightVal = midVal
			}
		}
	}
	return rightVal
}
