package problems

import "fmt"

// 32. Longest Valid Parentheses

func Problem_32() {
	s := "()()(())"
	fmt.Printf("res : %d\n", longestValidParentheses(s))
}

func longestValidParentheses(s string) int {
	max := 0
	stackCount := 0

	// Index : Stack count
	// Value : last index value of stack from which
	// we can take difference to calculate longest valid string
	windowArr := []int{-1}

	for i := 0; i <= len(s)-1; i++ {
		if s[i] == ')' {
			if stackCount == 0 {
				windowArr[0] = i
			} else {
				stackCount--
				val := i - windowArr[stackCount]
				if val > max {
					max = val
				}
				windowArr = windowArr[:stackCount+1]
			}
		} else {
			stackCount++
			windowArr = append(windowArr, i)
		}
	}

	return max
}
