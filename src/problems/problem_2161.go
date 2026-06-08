// 2161. Partition Array According to Given Pivot

package problems

import "fmt"

func Problem_2161() {
	nums := []int{9, 12, 5, 10, 14, 3, 10}
	pivot := 10
	fmt.Println(pivotArray(nums, pivot))
}

func pivotArray(nums []int, pivot int) []int {
	n := len(nums)
	res := make([]int, n)
	smallCount := 0
	largeCount := n - 1

	for i := 0; i < n; i++ {
		smallVal := nums[i]
		largeVal := nums[n-1-i]
		if smallVal < pivot {
			res[smallCount] = smallVal
			smallCount++
		}
		if largeVal > pivot {
			res[largeCount] = largeVal
			largeCount--
		}
	}

	for smallCount <= largeCount {
		res[smallCount] = pivot
		smallCount++
	}

	return res
}
