// 3471. Find the Largest Almost Missing Integer

package problems

import (
	"fmt"
	"slices"
)

func Problem_3471() {
	nums := []int{3, 9, 2, 1, 7}
	k := 3
	fmt.Println(largestInteger(nums, k))
}

// Constraints:

// ==> 1 <= nums.length <= 50
// ==> 0 <= nums[i] <= 50
// ==> 1 <= k <= nums.length

func largestInteger(nums []int, k int) int {
	n := len(nums)
	if k == n {
		slices.Max(nums)
	}
	count := make([]int, 51)
	for i := 0; i < n; i++ {
		count[nums[i]]++
	}
	if k > 1 {
		maxNum, minNum := nums[0], nums[n-1]
		if maxNum < minNum {
			maxNum, minNum = minNum, maxNum
		}
		if count[maxNum] == 1 {
			return maxNum
		}
		if count[minNum] == 1 {
			return minNum
		}
		return -1
	}

	for i := 50; i >= 0; i-- {
		if count[i] == 1 {
			return i
		}
	}

	return -1
}
