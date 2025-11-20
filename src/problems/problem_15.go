package problems

import (
	"fmt"
	"sort"
)

func Problem_15() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := [][]int{}
	for i := 0; i <= n-1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		first := nums[i]
		target := -first
		left := i + 1
		right := n - 1
		for left < right {
			if nums[left]+nums[right] == target {
				res = append(res, []int{first, nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if nums[left]+nums[right] < target {
				left++
			} else {
				right--
			}
		}

	}
	return res
}
