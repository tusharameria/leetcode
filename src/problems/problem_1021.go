// 1021. Remove Outermost Parentheses

package problems

import "fmt"

func Problem_1021() {
	s := "(())()((()))"
	fmt.Println(removeOuterParentheses(s))
}

func removeOuterParentheses(s string) string {
	count := 0
	sum := 0
	sLen := len(s)
	res := make([]byte, sLen)

	for i := 0; i < sLen; i++ {
		ch := s[i]
		if ch == '(' {
			sum++
			if sum != 1 {
				res[count] = ch
				count++
			}
		} else {
			sum--
			if sum != 0 {
				res[count] = ch
				count++
			}
		}

	}

	return string(res[:count])
}
