//3517. Smallest Palindromic Rearrangement I

package problems

import (
	"fmt"
)

func Problem_3517() {
	s := "babab"
	fmt.Println(smallestPalindrome(s))
}

func smallestPalindrome(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	res := make([]byte, n)
	if n%2 == 1 {
		res[n/2] = s[n/2]
	}

	countingStore := make([]uint16, 26)
	for i := 0; i < n/2; i++ {
		ch := s[i]
		countingStore[ch-'a']++
	}

	i := 0
	for idx, count := range countingStore {
		for range count {
			val := byte('a' + idx)
			res[i], res[n-1-i] = val, val
			i++
		}
	}

	return string(res)
}
