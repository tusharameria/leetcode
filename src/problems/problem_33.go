package problems

import "fmt"

// 33. Search in Rotated Sorted Array

func Problem_33() {
	nums := []int{1}
	target := 0
	fmt.Println("res : ", search(nums, target))
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		if nums[left] < nums[right] {
			if target < nums[left] || target > nums[right] {
				return -1
			}
		}
		if target < nums[left] && target > nums[right] {
			return -1
		}
		if target == nums[left] {
			return left
		}
		if target == nums[right] {
			return right
		}

		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > nums[right] {
			if target < nums[mid] && target > nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < nums[left] {
			if target > nums[mid] && target < nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}
