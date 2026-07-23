// 3513. Number of Unique XOR Triplets I

package problems

import (
	"fmt"
	"math/bits"
)

func Problem_3513() {
	nums := []int{3, 1, 2}
	fmt.Println(uniqueXorTriplets(nums))
}

func uniqueXorTriplets(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	return 1 << bits.Len(uint(n))
}
