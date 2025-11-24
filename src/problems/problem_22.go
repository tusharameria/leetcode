package problems

import "fmt"

func Problem_22() {
	n := 3
	fmt.Println(generateParenthesis(n))
}

func generateParenthesis(n int) []string {
	res := []string{}
	var solve func(open, close int, str string)

	solve = func(open, close int, str string) {
		if open > 0 {
			solve(open-1, close, str+"(")
		}
		if close > open {
			solve(open, close-1, str+")")
		}
		if len(str) == 2*n {
			res = append(res, str)
			return
		}
	}

	solve(n, n, "")

	return res
}
