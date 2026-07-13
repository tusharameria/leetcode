// 1331. Rank Transform of an Array

package problems

import (
	"fmt"
	"sort"
)

func Problem_1331() {
	arr := []int{40, 10, 20, 30}
	fmt.Println(arrayRankTransform(arr))
}

func arrayRankTransform(arr []int) []int {
	n := len(arr)
	sorted := make([]int, n)
	copy(sorted, arr)
	sort.Ints(sorted)

	store := make(map[int]int, n)
	r := 0
	for _, val := range sorted {
		if _, ok := store[val]; !ok {
			r++
			store[val] = r
		}
	}
	for i, val := range arr {
		sorted[i] = store[val]
	}
	return sorted
}
