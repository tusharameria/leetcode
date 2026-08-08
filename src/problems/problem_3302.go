// 3302. Find the Lexicographically Smallest Valid Sequence

package problems

import "fmt"

func Problem_3302() {
	word1 := "bacdc"
	word2 := "abc"
	fmt.Println(validSequence(word1, word2))
}

var suf [300_001]int32
var res [300_000]int

func validSequence(a string, b string) []int {
	m, n := int32(len(a)), int32(len(b))
	for i := range n {
		suf[i] = -1
	}
	suf[n] = m
	for i, j := m-1, n-1; i >= 0 && j >= 0; i-- {
		if a[i] == b[j] {
			suf[j] = i
			j--
		}
	}
	flag := false
	for i, j := int32(0), int32(0); i < m; i++ {
		if a[i] == b[j] {
			res[j] = int(i)
			j++
		} else if !flag && i < suf[j+1] {
			flag = true
			res[j] = int(i)
			j++
		}
		if j == n {
			return res[:j]
		}
	}
	return nil
}
