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

	for _, num := range nums {
		xor ^= num

		if num != 0 {
			hasNonZero = true
		}
	}

	if xor != 0 {
		return len(nums)
	}

	if hasNonZero {
		return len(nums) - 1
	}

	return 0
}
