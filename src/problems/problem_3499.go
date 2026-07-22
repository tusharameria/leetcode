// 3499. Maximize Active Section with Trade I

package problems

import "fmt"

func Problem_3499() {
	s := "1010101010100010"
	fmt.Println(maxActiveSectionsAfterTrade(s))
}

func maxActiveSectionsAfterTrade(s string) int {
	var i, ones, gain, prev int
	n := len(s)

	for ; i < n && s[i] == '1'; i++ {
		ones++
	}
	for ; i < n && s[i] == '0'; i++ {
		prev++
	}
	for i < n {
		for ; i < n && s[i] == '1'; i++ {
			ones++
		}
		if i == n {
			break
		}
		curr := 0
		for ; i < n && s[i] == '0'; i++ {
			curr++
		}
		gain = max(gain, prev+curr)
		prev = curr
	}

	return ones + gain
}
