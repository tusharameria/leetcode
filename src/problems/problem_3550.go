// 3550. Smallest Index With Digit Sum Equal to Index

package problems

import "fmt"

func Problem_3550() {
	nums := []int{3, 5, 1, 2, 3, 67, 23, 45, 12}
	fmt.Println(smallestIndex(nums))
}

func smallestIndex(nums []int) int {
	res := -1
	for i := len(nums) - 1; i >= 0; i-- {
		sum := 0
		num := nums[i]
		for num > 0 {
			sum += num % 10
			num /= 10
		}
		if sum == i {
			res = i
		}
	}
	return res
}
