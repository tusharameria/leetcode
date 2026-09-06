// 3904. Smallest Stable Index II

package problems

import (
	"fmt"
)

func Problem_3904() {
	fmt.Println(firstStableIndex([]int{}, 3))
}

// Constraints:

// --> 1 <= nums.length <= 10^5
// --> 0 <= nums[i] <= 10^9
// --> 0 <= k <= 10^9

func firstStableIndex(nums []int, k int) int {
	// fmt.Println("maxInt bit len", bits.Len(math.MaxInt))
	// fmt.Println("maxInt32 bit len", bits.Len(math.MaxInt32))
	// fmt.Println("maxInt bit len", bits.Len(math.MaxInt+1))
	// fmt.Println("10^9 bit len", bits.Len(1_000_000_000))
	n := len(nums)
	if n == 0 {
		return -1
	}
	resIdx := n
	maxArr := make([]int32, n)
	maxArr[0] = int32(nums[0])
	for i := 1; i < n; i++ {
		maxArr[i] = max(maxArr[i-1], int32(nums[i]))
	}
	minVal := int32(nums[n-1])
	if maxArr[n-1]-minVal <= int32(k) {
		resIdx = n - 1
	}
	for i := n - 2; i >= 0; i-- {
		minVal = min(minVal, int32(nums[i]))
		if maxArr[i]-minVal <= int32(k) {
			resIdx = i
		}
	}
	if resIdx == n {
		return -1
	}
	return resIdx
}
