package problems

import "fmt"

func Problem_5() {
	s := "cbbd"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	start, end := 0, 1
	for center := 0; center <= len(s)-1; center++ {
		l1, r1 := expandAroundCenter(s, center, center)
		if r1-l1 > end-start {
			start, end = l1, r1
		}
		l2, r2 := expandAroundCenter(s, center, center+1)
		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}
	return s[start:end]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for left >= 0 && right <= len(s)-1 && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right
}
