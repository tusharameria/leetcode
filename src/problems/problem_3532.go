// 3532. Path Existence Queries in a Graph I

package problems

import (
	"fmt"
)

func Problem_3532() {
	n := 4
	nums := []int{2, 5, 6, 8}
	maxDiff := 2
	queries := [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}
	fmt.Println(pathExistenceQueriesOld(n, nums, maxDiff, queries))
}

func pathExistenceQueriesOld(n int, nums []int, maxDiff int, queries [][]int) []bool {
	queryLen := len(queries)
	res := make([]bool, queryLen)
	if n == 0 {
		return res
	}
	store := make([]uint32, n)
	counter := uint32(0)
	prev := nums[0]
	for i := 1; i < n; i++ {
		curr := nums[i]
		if curr-prev > maxDiff {
			counter++
		}
		prev = curr
		store[i] = counter
	}

	for i := 0; i < queryLen; i++ {
		left, right := queries[i][0], queries[i][1]
		res[i] = store[left] == store[right]
	}

	return res
}
