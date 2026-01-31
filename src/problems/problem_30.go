package problems

import "fmt"

// 30. Substring with Concatenation of All Words

func Problem_30() {
	s := "barfootheabcbarfoo"
	words := []string{"foo", "bar", "the", "abc"}
	fmt.Println(findSubstring(s, words))
}

func findSubstring(s string, words []string) []int {

	wordLen := len(words)
	singleWordLen := len(words[0])
	totalWordLen := wordLen * singleWordLen
	res := []int{}

	for i := 0; i <= len(s)-totalWordLen; i++ {
		// fmt.Printf("=========== i : %d ===========\n", i)
		wordsCount := make(map[string]int)
		for _, word := range words {
			wordsCount[word]++
		}
		// fmt.Println("wordsCount : ", wordsCount)
		found := true
		for j := i; j-i <= totalWordLen-1; j += singleWordLen {
			// fmt.Printf("===== j : %d =====\n", j)
			targetStr := s[j : j+singleWordLen]
			// fmt.Printf("targetStr : %s\n", targetStr)
			if j+singleWordLen > len(s) {
				found = false
				break
			}
			if count, ok := wordsCount[targetStr]; !ok || count <= 0 {
				found = false
				break
			}
			wordsCount[targetStr]--
		}
		for _, val := range wordsCount {
			// fmt.Printf("key : %s, val : %d\n", key, val)
			if val != 0 {
				found = false
				break
			}
		}
		if found {
			// fmt.Println("found i : ", i)
			res = append(res, i)
		}
	}
	return res
}
