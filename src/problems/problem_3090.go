// 3090. Maximum Length Substring With Two Occurrences

package problems

import "fmt"

func Problem_3090() {
	s := "bcbbbcba"
	fmt.Println(maximumLengthSubstring(s))
}

// Constraints

// ==> 2 <= s.length <= 100
// ==> s consists only of lowercase English letters.

func maximumLengthSubstring(s string) int {
	n := len(s)
	maxLen := 1
	store := make([]uint8, 26)
	left := 0
	right := 0
	for i := right; i < n; i++ {
		idx := int(s[i] - 'a')

		if store[idx] == 2 {
			maxLen = max(maxLen, i-left)
			for j := left; j < i; j++ {
				newIdx := int(s[j] - 'a')
				store[newIdx]--
				if newIdx == idx {
					left = j + 1
					break
				}
			}
		}
		store[idx]++
	}

	return max(maxLen, n-left)
}
