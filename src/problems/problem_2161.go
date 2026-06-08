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
	smallRes := make([]int, n)
	sameRes := make([]int, n)
	largeRes := make([]int, n)
	smallCount, sameCount, largeCount := 0, 0, 0

	for _, val := range nums {
		if val < pivot {
			smallRes[smallCount] = val
			smallCount++
		} else if val == pivot {
			sameRes[sameCount] = val
			sameCount++
		} else {
			largeRes[largeCount] = val
			largeCount++
		}
	}
	return append(smallRes[:smallCount], append(sameRes[:sameCount], largeRes[:largeCount]...)...)
}
