package problems

import "fmt"

func Problem_3() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	maxLength := 1
	bytesIndex := make(map[byte]int)
	i := 0
	j := 1
	bytesIndex[s[0]] = 0
	for i <= len(s)-1 && j <= len(s)-1 {
		if _, ok := bytesIndex[s[j]]; !ok {
			maxLength = max(maxLength, j-i+1)
		} else {
			for k := bytesIndex[s[j]] - 1; k >= i; k-- {
				delete(bytesIndex, s[k])
			}
			i = bytesIndex[s[j]] + 1
		}
		bytesIndex[s[j]] = j
		j++
	}
	return maxLength
}
