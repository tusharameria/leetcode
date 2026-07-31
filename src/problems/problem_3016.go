// 3016. Minimum Number of Pushes to Type Word II

package problems

import (
	"fmt"
	"sort"
)

func Problem_3016() {
	word := "aabbccddeeffgghhiiiiii"
	fmt.Println(minimumPushes(word))
}

func minimumPushes(word string) int {
	count := make([]int, 26)
	for _, ch := range word {
		count[ch-'a']++
	}
	sort.Ints(count)

	res := 0
	for i := 25; i >= 0 && count[i] != 0; i-- {
		res += (((25 - i) / 8) + 1) * count[i]
	}
	return res
}
