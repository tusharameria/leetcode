// 2213. Longest Substring of One Repeating Character

package problems

import (
	"fmt"
)

func Problem_2213() {
	s := "jdbdhbfsjbjhbhjjsjdb"
	queryCharacters := "ddd"
	queryIndices := []int{2, 4, 5}
	fmt.Println(longestRepeating(s, queryCharacters, queryIndices))
}

func longestRepeating(s string, queryCharacters string, queryIndices []int) []int {
	b := []byte(s)
	n := len(queryIndices)
	res := make([]int, n)
	for i := range n {
		idx := queryIndices[i]
		b[idx] = queryCharacters[i]
		res[i] = int(longest_ch_len_2213(b))
	}
	return res
}

func longest_ch_len_2213(b []byte) uint32 {
	var maxVal uint32 = 1
	var currentStreak uint32 = 1
	for i := 1; i < len(b); i++ {
		if b[i] == b[i-1] {
			currentStreak++
		} else {
			maxVal = max(maxVal, currentStreak)
			currentStreak = 1
		}
	}
	return max(maxVal, currentStreak)
}
