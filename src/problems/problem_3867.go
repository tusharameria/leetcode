// 3867. Sum of GCD of Formed Pairs

package problems

import (
	"fmt"
	"sort"
)

func Problem_3867() {
	nums := []int{3, 6, 2, 8}
	fmt.Println(nums)
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
		prefixGcd[i] = gcd(maxVal, currVal)
	}

	res := int64(0)
	sort.Ints(prefixGcd)
	for i := 0; i < (n / 2); i++ {
		res += int64(gcd(prefixGcd[i], prefixGcd[n-1-i]))
	}
	return res
}

func gcd(a, b int) int {
	if b < a {
		a, b = b, a
	}

	for b%a != 0 {
		a, b = b%a, a
	}

	return a
}
