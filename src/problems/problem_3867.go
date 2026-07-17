// 3867. Sum of GCD of Formed Pairs

package problems

import (
	"fmt"
	"sort"
)

func Problem_3867() {
	nums := []int{1, 19, 33, 39, 40, 53, 65, 62}
	fmt.Println(gcdSum(nums))
}

const SIZE = 64

var LUT [SIZE][SIZE]uint8

func init() {
	for i := uint8(1); i < SIZE; i++ {
		LUT[0][i], LUT[i][0], LUT[i][i] = i, i, i
		for j := i + 1; j < SIZE; j++ {
			r := LUT[i][j-i]
			LUT[i][j], LUT[j][i] = r, r
		}
	}
}

func gcdSum(nums []int) int64 {
	n := len(nums)
	if n == 1 {
		return 0
	}
	prefixGcd := make([]int, n)
	maxVal := nums[0]
	prefixGcd[0] = maxVal
	for i := 1; i < n; i++ {
		currVal := nums[i]
		maxVal = max(maxVal, currVal)
		prefixGcd[i] = gcd(currVal, maxVal)
	}

	fmt.Println(prefixGcd)

	res := int64(0)
	sort.Ints(prefixGcd)
	for i := 0; i < (n / 2); i++ {
		res += int64(gcd(prefixGcd[i], prefixGcd[n-1-i]))
	}
	return res
}

// a is smaller
func gcd(a, b int) int {
	for a >= SIZE {
		a, b = b%a, a
	}
	if a == 0 {
		return b
	}
	return int(LUT[b%a][a])
}
