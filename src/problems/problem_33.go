package problems

import "fmt"

// 33. Search in Rotated Sorted Array

func Problem_33() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 1
	fmt.Println("res : ", search(nums, target))
}

func search(nums []int, target int) int {
	n := len(nums)
	if n == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	left, right := 0, n-1

	for left < right {
		leftVal := nums[left]
		rightVal := nums[right]
		if leftVal == target {
			return left
		}
		if rightVal == target {
			return right
		}
		mid := left + (right-left)/2
		midVal := nums[mid]
		if midVal == target {
			return mid
		}

		if leftVal < rightVal {
			if target < midVal {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if leftVal < midVal {
				// rightVal < leftVal < midVal
				if target > midVal {
					// temporary
					left = mid + 1
				} else if target > leftVal {
					right = mid - 1
				} else if target > rightVal {
					return -1
				} else {
					left = mid + 1
				}
			} else {
				//  midVal < rightVal < leftVal
				if target > leftVal {
					right = mid - 1
				} else if target > rightVal {
					return -1
				} else if target > midVal {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
		}
	}

	return -1
}
