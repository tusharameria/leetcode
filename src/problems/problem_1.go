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
	type NumInfo struct {
		num int
		idx int
	}
	n := len(nums)
	numsInfo := make([]NumInfo, n)
	for i, val := range nums {
		numsInfo[i] = NumInfo{
			num: val,
			idx: i,
		}
	}
	sort.Slice(numsInfo, func(i, j int) bool {
		return numsInfo[i].num < numsInfo[j].num
	})
	i, j := 0, len(nums)-1
	for i < j {
		leftVal := numsInfo[i].num
		rightVal := numsInfo[j].num
		if leftVal+rightVal == target {
			return []int{numsInfo[i].idx, numsInfo[j].idx}
		} else if leftVal+rightVal < target {
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

func twoSumMap(nums []int, target int) []int {
	n := len(nums)
	store := make(map[int]int, n)
	for i, val := range nums {
		if idx, ok := store[target-val]; !ok {
			store[val] = i
		} else {
			return []int{i, idx}
		}
	}
	return nil
}
