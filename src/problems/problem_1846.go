// 1846. Maximum Element After Decreasing and Rearranging

package problems

import (
	"fmt"
	"sort"
)

func Problem_1846() {
	nums := []int{2, 2, 1, 2, 1}
	fmt.Println(maximumElementAfterDecrementingAndRearranging(nums))
}

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)
	n := len(arr)
	arr[0] = 1
	for i := 1; i < n; i++ {
		target := arr[i-1] + 1
		if arr[i] > target {
			arr[i] = target
		}
	}
	return arr[n-1]
}
