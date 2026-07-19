// 1081. Smallest Subsequence of Distinct Characters

package problems

import "fmt"

func Problem_1081() {
	s := "bcabc"
	fmt.Println(smallestSubsequence(s))
}

func smallestSubsequence(s string) string {
	freq := make([]uint16, 26)
	for _, ch := range s {
		freq[ch-'a']++
	}
	inStack := make([]bool, 26)
	stack := make([]byte, 0)

	for i := range s {
		ch := s[i]
		freq[ch-'a']--
		if !inStack[ch-'a'] {
			for len(stack) > 0 && freq[stack[len(stack)-1]-'a'] > 0 && stack[len(stack)-1] > ch {
				inStack[stack[len(stack)-1]-'a'] = false
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, ch)
			inStack[ch-'a'] = true
		}
	}

	return string(stack)
}
