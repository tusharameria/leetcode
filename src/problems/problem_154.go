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
	var findMinIter func(left, right, leftVal, rightVal int) int
	findMinIter = func(left, right, leftVal, rightVal int) int {
		for left < right {
			if right-left == 1 {
				return min(leftVal, rightVal)
			}
			mid := left + (right-left)/2
			midVal := nums[mid]
			if leftVal < rightVal {
				return leftVal
			} else {
				if leftVal == rightVal && leftVal == midVal {
					return min(findMinIter(left, mid-1, leftVal, nums[mid-1]), findMinIter(mid+1, right, nums[mid+1], rightVal))
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
		}
		return rightVal
	}
	n := len(nums)
	left := 0
	right := n - 1
	leftVal := nums[left]
	rightVal := nums[right]

	return findMinIter(left, right, leftVal, rightVal)
}
