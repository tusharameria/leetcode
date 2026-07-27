// 628. Maximum Product of Three Numbers

package problems

import (
	"fmt"
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
	max1, max2, max3 := -1000, -1000, -1000
	min1, min2 := 1000, 1000

	for i := 0; i < n; i++ {
		val := nums[i]
		if val > max1 {
			max1, max2, max3 = val, max1, max2
		} else if val > max2 {
			max2, max3 = val, max2
		} else if val > max3 {
			max3 = val
		}

		if val < min1 {
			min1, min2 = val, min1
		} else if val < min2 {
			min2 = val
		}

	}

	return max((max1 * max2 * max3), (max1 * min1 * min2))
}
