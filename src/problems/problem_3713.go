package problems

import "fmt"

// 3713. Longest Balanced Substring I

func Problem_3713() {
	s := "aabbccccdddd"
	fmt.Println("res : ", longestBalanced(s))
}

func longestBalanced(s string) int {
	res := 1
	lens := len(s)
	for i := 0; i <= lens-1-res; i++ {
		letterCounts := make(map[byte]int)
		letterCounts[s[i]]++
		for j := i + 1; j <= lens-1; j++ {
			letterCounts[s[j]]++
			val := letterCounts[s[j]]
			allEqual := true
			for _, count := range letterCounts {
				if count != val {
					allEqual = false
					break
				}
			}
			if allEqual && j-i+1 > res {
				res = j - i + 1
			}
		}
	}
	return res
}
