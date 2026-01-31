package problems

import "fmt"

// 30. Substring with Concatenation of All Words

func Problem_30() {
	s := "barfootheabcbarfoofoobartheabcfo"
	words := []string{"foo", "bar", "the", "abc"}
	fmt.Println(findSubstring(s, words))
}

func findSubstring(s string, words []string) []int {

	wordLen := len(words)
	singleWordLen := len(words[0])
	totalWordLen := wordLen * singleWordLen
	res := []int{}

	orgFreqMap := make(map[string]int)
	for _, word := range words {
		orgFreqMap[word]++
	}

	for i := 0; i <= singleWordLen-1 && i+totalWordLen <= len(s); i++ {
		// Initial required values for each word
		reqFreqMap := make(map[string]int)
		for _, word := range words {
			reqFreqMap[word]++
		}

		// Values for first iteration
		for j := 0; j <= wordLen-1; j++ {
			if (j+1)*singleWordLen+i > len(s) {
				break
			}
			testStr := s[j*singleWordLen+i : (j+1)*singleWordLen+i]
			if _, ok := orgFreqMap[testStr]; ok {
				reqFreqMap[testStr]--
				if reqFreqMap[testStr] == 0 {
					delete(reqFreqMap, testStr)
				}
			}
		}

		if len(reqFreqMap) == 0 {
			res = append(res, i)
		}

		initSubStr := s[i : i+singleWordLen]
		if _, ok := orgFreqMap[initSubStr]; ok {
			reqFreqMap[initSubStr]++
			if reqFreqMap[initSubStr] == 0 {
				delete(reqFreqMap, initSubStr)
			}
		}

		for j := i + totalWordLen; j <= len(s)-1; j += singleWordLen {
			if j+singleWordLen > len(s) {
				break
			}
			newSubStr := s[j : j+singleWordLen]
			if _, ok := orgFreqMap[newSubStr]; ok {
				reqFreqMap[newSubStr]--
				if reqFreqMap[newSubStr] == 0 {
					delete(reqFreqMap, newSubStr)
				}
			}

			if len(reqFreqMap) == 0 {
				res = append(res, j-totalWordLen+singleWordLen)
			}

			oldSubStr := s[j-totalWordLen+singleWordLen : j-totalWordLen+singleWordLen+singleWordLen]
			if _, ok := orgFreqMap[oldSubStr]; ok {
				reqFreqMap[oldSubStr]++
				if reqFreqMap[oldSubStr] == 0 {
					delete(reqFreqMap, oldSubStr)
				}
			}
		}

	}

	return res
}

func findSubstringOld(s string, words []string) []int {

	wordLen := len(words)
	singleWordLen := len(words[0])
	totalWordLen := wordLen * singleWordLen
	res := []int{}

	for i := 0; i <= len(s)-totalWordLen; {
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
