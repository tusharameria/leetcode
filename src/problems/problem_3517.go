//3517. Smallest Palindromic Rearrangement I

package problems

import (
	"fmt"
	"slices"
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
	buff := make([]byte, n/2)
	copy(buff, s[:n/2])

	slices.Sort(buff)

	for i := 0; i < n/2; i++ {
		val := buff[i]
		res[i], res[n-1-i] = val, val
	}

	return string(res)
}
