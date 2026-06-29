// 1967. Number of Strings That Appear as Substrings in Word

package problems

import (
	"fmt"
	"strings"
)

func Problem_1967() {
	patterns := []string{"a", "abc", "bc", "d"}
	word := "abc"
	fmt.Println(numOfStrings(patterns, word))
}

func numOfStrings(patterns []string, word string) int {
	res := 0
	for _, str := range patterns {
		if strings.Contains(word, str) {
			res++
		}
	}
	return res
}
