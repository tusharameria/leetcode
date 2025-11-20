package problems

import (
	"fmt"
	"sort"
)

func Problem_1() {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Printf("%v\n", twoSumEff(nums, target))
}

func twoSumBrute(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumEff(nums []int, target int) []int {
	numsBuff := []int{}
	for i := 0; i <= len(nums)-1; i++ {
		numsBuff = append(numsBuff, nums[i])
	}
	sort.Ints(nums)
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]+nums[j] == target {
			res := []int{}
			fmt.Println(numsBuff)
			for a := 0; a <= len(numsBuff)-1; a++ {
				if nums[i] == numsBuff[a] || nums[j] == numsBuff[a] {
					res = append(res, a)
				}
			}
			return res
		}
		if nums[i]+nums[j] < target {
			i++
		} else {
			j--
		}
	}
	return nil
}

func twoSumBin(nums []int, target int) []int {
	sort.Ints(nums)

	for i := 0; i <= len(nums)-1; i++ {
		if index := BinarySearchInt(nums[i+1:], target-nums[i]); index != -1 {
			return []int{i, i + index}
		}
	}
	return []int{}
}
