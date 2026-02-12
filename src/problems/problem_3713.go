package problems

import "fmt"

// 3713. Longest Balanced Substring I

func Problem_3713() {
	s := "aabbccccdddd"
	fmt.Println("res : ", longestBalanced(s))
}

func longestBalanced(s string) int {
	res := 0
	lens := len(s)
	for i := 0; i <= lens-1-res; i++ {
		letterFreq := make([]int, 26)
		distinctCount := 0
		maxFreq := 0
		elementsWithMaxFreq := 0
		for j := i; j <= lens-1; j++ {
			index := int(s[j] - 'a')
			letterFreq[index]++
			if letterFreq[index] == 1 {
				distinctCount++
			}
			if letterFreq[index] == maxFreq {
				elementsWithMaxFreq++
			}
			if letterFreq[index] > maxFreq {
				maxFreq = letterFreq[index]
				elementsWithMaxFreq = 1
			}
			if distinctCount == elementsWithMaxFreq && j-i+1 > res {
				res = j - i + 1
			}
		}
	}
	return res
}
