package problems

import "fmt"

func Problem_20() {
	s := "(({[}]))"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	stack := []byte{}
	for i := 0; i <= len(s)-1; i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			length := len(stack)
			if length == 0 || (s[i] == ')' && stack[length-1] != '(') || (s[i] == '}' && stack[length-1] != '{') || (s[i] == ']' && stack[length-1] != '[') {
				return false
			} else {
				stack = stack[:length-1]
			}
		}
	}
	return len(stack) == 0
}
