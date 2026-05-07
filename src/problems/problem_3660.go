// 3660. Jump Game IX

package problems

import "fmt"

func Problem_3660() {
	nums := []int{9, 30, 16, 6, 36, 9}
	fmt.Println(maxValue(nums))
}

func maxValue(nums []int) []int {
	n := len(nums)
	mmax := 0
	mmin := int(^uint(0) >> 1)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		if nums[i] > mmax {
			mmax = nums[i]
		}
		ans[i] = mmax
	}
	for i := n - 1; i >= 0; i-- {
		if ans[i] > mmin {
			ans[i] = ans[i+1]
		}
		if nums[i] < mmin {
			mmin = nums[i]
		}
	}
	return ans
}
