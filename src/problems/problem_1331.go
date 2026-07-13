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
	pos := make([]int, n)
	for i := range n {
		pos[i] = i
	}

	sort.Slice(pos, func(i, j int) bool {
		return arr[pos[i]] < arr[pos[j]]
	})

	fmt.Println(pos)

	res := make([]int, n)
	val := 1
	res[pos[0]] = val
	for i := 1; i < n; i++ {
		if arr[pos[i]] != arr[pos[i-1]] {
			val++
		}
		res[pos[i]] = val
	}
	return res
}
