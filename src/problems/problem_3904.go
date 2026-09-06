// 3904. Smallest Stable Index II

package problems

import (
	"fmt"
)

func Problem_3904() {
	fmt.Println(firstStableIndex([]int{5, 0, 1, 4}, 3))
}

// Constraints:

// --> 1 <= nums.length <= 10^5
// --> 0 <= nums[i] <= 10^9
// --> 0 <= k <= 10^9

const MASK_3904 = (1 << 32) - 1

func firstStableIndex(nums []int, k int) int {
	n := len(nums)

	// Max Values from 0 to i in uppar bits
	maxVal := -1
	for i := 0; i < n; i++ {
		maxVal = max(maxVal, nums[i])
		nums[i] |= maxVal << 32
	}

	resIdx := n
	minVal := 1_000_000_001
	for i := n - 1; i >= 0; i-- {
		lower := nums[i] & MASK_3904
		minVal = min(minVal, lower)
		upper := nums[i] >> 32
		if upper-minVal <= k {
			resIdx = i
		}
	}

	if resIdx == n {
		return -1
	}
	return resIdx
}
