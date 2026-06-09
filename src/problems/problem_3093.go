// 3093. Longest Common Suffix Queries

package problems

import "fmt"

func Problem_3093() {
	wordsContainer := []string{"abcdefgh", "poiuygh", "ghghgh"}
	wordsQuery := []string{"gh", "acbfgh", "acbfegh"}
	fmt.Println(stringIndices(wordsContainer, wordsQuery))
}

func stringIndices(wordsContainer []string, wordsQuery []string) []int {
	len2 := len(wordsQuery)
	res := make([]int, len2)
	minLen := 10_000
	minLenIndex := 0
	for i, str := range wordsContainer {
		if len(str) < minLen {
			minLen = len(str)
			minLenIndex = i
		}
	}

	len1 := len(wordsContainer)

	for i := 0; i < len2; i++ {
		maxVal := 0
		maxValLen := 0
		maxLenIndex := 0
		for j := 0; j < len1; j++ {
			val := longestSuffixLen(wordsQuery[i], wordsContainer[j])
			if val != 0 {
				if val >= maxVal {
					if val > maxVal {
						maxVal = val
						maxValLen = len(wordsContainer[j])
						maxLenIndex = j
					} else {
						if len(wordsContainer[j]) < maxValLen {
							maxValLen = len(wordsContainer[j])
							maxLenIndex = j
						}
					}
				}
			}
		}
		if maxVal == 0 {
			res[i] = minLenIndex
		} else {
			res[i] = maxLenIndex
		}
	}

	return res
}

func longestSuffixLen(query, container string) int {
	queryLen := len(query)
	containeLen := len(container)
	len := 0
	for i, j := queryLen-1, containeLen-1; i >= 0 && j >= 0; {
		if query[i] == container[j] {
			len++
		} else {
			break
		}
		i--
		j--
	}
	return len
}
