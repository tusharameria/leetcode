// 628. Maximum Product of Three Numbers

package problems

import (
	"fmt"
	"sort"
)

func Problem_628() {
	nums := []int{2, 4, 5, -1, -5, -9}
	fmt.Println(maximumProduct(nums))
}

func maximumProduct(nums []int) int {
	n := len(nums)
	if n == 3 {
		return nums[0] * nums[1] * nums[2]
	}
	sort.Ints(nums)
	return max((nums[n-1] * nums[n-2] * nums[n-3]), (nums[n-1] * nums[0] * nums[1]))
}
