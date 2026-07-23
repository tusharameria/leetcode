// 3513. Number of Unique XOR Triplets I

package problems

import (
	"fmt"
)

func Problem_3513() {
	nums := []int{3, 1, 2}
	fmt.Println(uniqueXorTriplets(nums))
}

func uniqueXorTriplets(nums []int) int {
	n := len(nums)
	res := make(map[int]bool)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			firstXor := nums[i] ^ nums[j]
			for k := j; k < n; k++ {
				res[firstXor^nums[k]] = true
			}
		}
	}

	return len(res)
}
