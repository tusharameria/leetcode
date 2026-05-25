// 1752. Check if Array Is Sorted and Rotated

package problems

import "fmt"

func Problem_1752() {
	nums := []int{3, 4, 4, 5, 1, 4}
	fmt.Println(check(nums))
}

func check(nums []int) bool {
	decreased := 0
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			decreased++
		}
		if decreased == 2 {
			return false
		}
	}
	if nums[0] < nums[n-1] {
		decreased++
	}
	return decreased <= 1
}
