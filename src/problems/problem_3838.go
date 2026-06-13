// 3838. Weighted Word Mapping

package problems

import "fmt"

func Problem_3838() {
	words := []string{"abcd", "def", "xyz"}
	weights := []int{5, 3, 12, 14, 1, 2, 3, 2, 10, 6, 6, 9, 7, 8, 7, 10, 8, 9, 6, 9, 9, 8, 3, 7, 7, 2}
	fmt.Println(mapWordWeights(words, weights))
}

func mapWordWeights(words []string, weights []int) string {
	n := len(words)
	res := make([]byte, n)

	for i, word := range words {
		sum := 0
		for _, letter := range word {
			sum += weights[letter-'a']
		}
		res[i] = uint8('a') + uint8(25-sum%26)
	}

	return string(res)
}
