package problems

import "fmt"

func Problem_1930() {
	s := "bbcbaba"
	fmt.Println(countPalindromicSubsequence(s))
}

func countPalindromicSubsequence(s string) int {
	hash := make(map[byte]bool)
	count := 0
	for i := 0; i <= len(s)-3; i++ {
		if _, ok := hash[s[i]]; !ok {
			j := len(s) - 1
			for j > i+1 && s[i] != s[j] {
				j--
			}
			if j == i+1 {
				hash[s[i]] = true
				continue
			}
			hashMid := make(map[byte]bool)
			for k := i + 1; k < j; k++ {
				hashMid[s[k]] = true
			}
			count += len(hashMid)
			hash[s[i]] = true
		}
	}
	return count
}

func countPalindromicSubsequenceEff(s string) int {
	n := len(s)
	first := make([]int, 26)
	last := make([]int, 26)
	result := 0

	// initialize first to a sentinel
	for i := range first {
		first[i] = -1
		last[i] = -1
	}

	for i := 0; i <= n-1; i++ {
		index := int(s[i] - 'a')
		if first[index] == -1 {
			first[index] = i
		}
		last[index] = i
	}
	for c := 0; c < 26; c++ {
		if first[c] != -1 && last[c] != -1 && first[c] < last[c] {
			mid := make(map[byte]bool)
			for k := first[c] + 1; k < last[c]; k++ {
				mid[s[k]] = true
			}
			result += len(mid)
		}
	}
	return result
}
