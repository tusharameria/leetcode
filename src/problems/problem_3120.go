// 3120. Count the Number of Special Characters I

package problems

import "fmt"

func Problem_3120() {
	word := "sdSaD"
	fmt.Println(numberOfSpecialCharsOld(word))
}

func numberOfSpecialCharsOld(word string) int {
	count := 0
	store := make([]bool, 64)

	for i := 0; i < len(word); i++ {
		store[int(word[i])-65] = true
	}

	for i := 0; i < 26; i++ {
		if store[i] && store[i+32] {
			count++
		}
	}

	return count
}
