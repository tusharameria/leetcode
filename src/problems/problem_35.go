package problems

import "fmt"

// 35. Search Insert Position

func Problem_35() {
	nums := []int{0, 1, 2, 3, 4, 5, 7}
	target := 6
	fmt.Println(searchInsert(nums, target))
}

func searchInsert(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	left := 0
	right := n - 1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if nums[left] < target {
		return left + 1
	}
	return left
}
