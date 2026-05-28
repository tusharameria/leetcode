// 3120. Count the Number of Special Characters II

package problems

import "fmt"

func Problem_3121() {
	word := "aaAbcBC"
	fmt.Println(numberOfSpecialChars(word))
}

func numberOfSpecialChars(word string) int {
	count := 0
	store := make([]bool, 64)
	check := make([]bool, 26)

	for i := 0; i < len(word); i++ {
		val := int(word[i])
		store[val-65] = true
		if val >= 97 {
			if store[val-65-32] {
				check[val-32-65] = true
			}
		}
	}

	for i := 0; i < 26; i++ {
		if store[i] && store[i+32] && !check[i] {
			count++
		}
	}

	return count
}
