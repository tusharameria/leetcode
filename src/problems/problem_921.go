// 921. Minimum Add to Make Parentheses Valid

package problems

import "fmt"

func Problem_921() {
	s := "))(()"
	fmt.Println(minAddToMakeValid(s))
}

func minAddToMakeValid(s string) int {
	res := 0
	sum := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			if sum < 0 {
				res -= sum
				sum = 0
			}
			sum++
		} else {
			sum--
		}
	}
	if sum < 0 {
		sum = -sum
	}
	return res + sum
}
