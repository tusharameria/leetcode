// 1464. Maximum Product of Two Elements in an Array

package problems

import "fmt"

func Problem_1464() {
	nums := []int{1, 4, 5}
	fmt.Println(maxProduct(nums))
}

func maxProduct(nums []int) int {
	max1, max2 := 0, 0
	for _, val := range nums {
		if val > max1 {
			max1, max2 = val, max1
		} else if val > max2 {
			max2 = val
		}
	}
	return (max1 - 1) * (max2 - 1)
}
