package problems

import (
	"fmt"
	"sort"
)

func Problem_18() {
	nums := []int{}
	target := 0
	fmt.Println(fourSum(nums, target))
}

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	a := len(nums)
	for i := 0; i <= a-4; i++ {
		for j := i + 1; j <= a-3; j++ {
			internalTarget := target - nums[i] - nums[j]
			left := j + 1
			right := a - 1
			if internalTarget >= nums[left]*2 && internalTarget <= nums[right]*2 {
				for left < right {
					if nums[left]+nums[right] == internalTarget {
						res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
						for left < right && nums[left] == nums[left+1] {
							left++
						}
						for left < right && nums[right] == nums[right-1] {
							right--
						}
						left++
						right--
					} else if nums[left]+nums[right] < internalTarget {
						left++
					} else {
						right--
					}
				}
			}
			for j <= a-3 && nums[j] == nums[j+1] {
				j++
			}
		}
		for i <= a-4 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}
