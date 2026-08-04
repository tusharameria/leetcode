// 3731. Find Missing Elements

package problems

import (
	"fmt"
	"slices"
)

func Problem_3731() {
	nums := []int{2, 9, 6, 3, 8}
	fmt.Println(findMissingElements(nums))
}

func findMissingElements(nums []int) []int {
	minVal := slices.Min(nums)
	maxVal := slices.Max(nums)
	n := len(nums)
	newLen := (maxVal - minVal + 1) - n
	if newLen == 0 {
		return []int{}
	}

	res := make([]int, newLen)
	store := make([]bool, maxVal-minVal+1)
	for _, val := range nums {
		store[val-minVal] = true
	}

	j := 0
	for i, pres := range store {
		if !pres {
			res[j] = i + minVal
			j++
		}
	}
	return res
}
