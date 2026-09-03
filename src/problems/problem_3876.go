// 3876. Construct Uniform Parity Array II

package problems

import (
	"fmt"
	"math"
)

func Problem_3876() {
	nums1 := []int{1, 4, 7}
	fmt.Println(uniformArray(nums1))
}

// Constraints:

// --> 1 <= n == nums1.length <= 10^5
// --> 1 <= nums1[i] <= 10^9
// --> nums1 consists of distinct integers.

const MAXINT = math.MaxInt

func uniformArray(nums1 []int) bool {
	minEven := MAXINT
	minOdd := MAXINT
	for _, val := range nums1 {
		if val%2 == 0 {
			minEven = min(val, minEven)
		} else {
			minOdd = min(val, minOdd)
		}
	}
	if minEven == MAXINT || minOdd == MAXINT {
		return true
	}

	return minEven > minOdd
}
