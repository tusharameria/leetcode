// 3702. Longest Subsequence With Non-Zero Bitwise XOR

package problems

import "fmt"

func Problem_3702() {
	nums := []int{1, 2, 3}
	fmt.Println(longestSubsequence(nums))
}

// Constraints:

// --> 1 <= nums.length <= 10^5
// --> 0 <= nums[i] <= 10^9

func longestSubsequence(nums []int) int {
	xor := 0
	hasNonZero := false
	n := len(nums)
	for i := 0; i < n; i++ {
		num := nums[i]
		if num != 0 {
			hasNonZero = true
			break
		}
	}

	if !hasNonZero {
		return 0
	}

	for i := 0; i < n; i++ {
		num := nums[i]
		xor ^= num
	}

	if xor != 0 {
		return n
	}

	return n - 1
}
