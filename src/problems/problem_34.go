package problems

import "fmt"

// 34. Find First and Last Position of Element in Sorted Array

func Problem_34() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println("res : ", searchRange(nums, target))
}

func searchRange(nums []int, target int) []int {
	res := -1
	left := 0
	right := len(nums) - 1
	mid := left + (right-left)/2

	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] == target {
			res = mid
			break
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if res == -1 {
		return []int{-1, -1}
	}
	return []int{searchNum(nums, left, mid, target, true), searchNum(nums, mid, right, target, false)}
}

func searchNum(nums []int, left, right, target int, findFirst bool) int {
	res := left
	if findFirst {
		res = right
	}
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			res = mid
			if findFirst {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}
