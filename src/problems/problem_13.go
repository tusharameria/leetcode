package problems

import (
	"fmt"
)

func Problem_13() {
	s := ""
	fmt.Println(romanToInt(s))
}

func romanToInt(s string) int {
	values := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	sum := 0
	n := len(s)
	for i := 0; i <= n-1; i++ {
		if i+1 < n && values[s[i]] < values[s[i+1]] {
			sum -= values[s[i]]
		} else {
			sum += values[s[i]]
		}
	}
	return sum
}
