// 3514. Number of Unique XOR Triplets II

package problems

import (
	"fmt"
)

func Problem_3514() {
	nums := []int{1280, 476, 379}
	fmt.Println(uniqueXorTriplets(nums))
}

func uniqueXorTriplets(nums []int) int {

	const MAXX = 2048

	pair := make([]bool, MAXX)
	triplet := make([]bool, MAXX)

	n := len(nums)

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			pair[nums[i]^nums[j]] = true
		}
	}

	for x := 0; x < MAXX; x++ {

		if !pair[x] {
			continue
		}

		for _, v := range nums {
			triplet[x^v] = true
		}
	}

	ans := 0

	for _, ok := range triplet {
		if ok {
			ans++
		}
	}

	return ans
}
