// 3016. Minimum Number of Pushes to Type Word II

package problems

import (
	"fmt"
)

func Problem_3016() {
	word := "abccdeggg"
	fmt.Println(minimumPushes(word))
}

func minimumPushes(word string) int {
	count := make([]uint32, 26)
	for _, ch := range word {
		count[ch-'a']++
	}
	sort_3016(count)

	res := 0
	for i, j := 0, 0; count[i] != 0; i++ {
		if i%8 == 0 {
			j++
		}
		res += int(count[i]) * j
	}
	return res
}

func sort_3016(arr []uint32) {
	for i := 1; i < 26; i++ {
		for j := i; j > 0 && arr[j] > arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}
