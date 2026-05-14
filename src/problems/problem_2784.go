// 2784. Check if Array is Good

package problems

import (
	"fmt"
)

func Problem_2784() {
	nums := []int{1, 3, 2, 5}
	fmt.Println(isGood(nums))
}

func isGood(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return false
	}
	count := make([]int, 200)
	for _, num := range nums {
		count[num-1]++
	}
	for i := 0; i < n-2; i++ {
		if count[i] != 1 {
			return false
		}
	}
	return count[n-2] == 2
}
