// 2553. Separate the Digits in an Array

package problems

import "fmt"

func Problem_2553() {
	nums := []int{0, 23, 45, 67}
	fmt.Println(separateDigits(nums))
}

func separateDigits(nums []int) []int {
	res := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		if num == 0 {
			res = append(res, 0)
			continue
		}
		for num > 0 {
			res = append(res, num%10)
			num /= 10
		}
	}
	n := len(res)
	for i := 0; i < n/2; i++ {
		res[i], res[n-1-i] = res[n-1-i], res[i]
	}
	return res
}
