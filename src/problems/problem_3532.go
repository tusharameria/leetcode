// 3532. Path Existence Queries in a Graph I

package problems

import "fmt"

func Problem_3532() {
	n := 4
	nums := []int{2, 5, 6, 8}
	maxDiff := 2
	queries := [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}
	fmt.Println(pathExistenceQueries(n, nums, maxDiff, queries))
}

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
	queryLen := len(queries)
	res := make([]bool, queryLen)
	store := make([]bool, n)
	store[0] = true

	for i := 1; i < n; i++ {
		store[i] = nums[i]-nums[i-1] <= maxDiff
	}

	for i := 0; i < queryLen; i++ {
		startNode, endNode := queries[i][0], queries[i][1]
		if startNode > endNode {
			startNode, endNode = endNode, startNode
		}
		val := true
		if startNode != endNode {
			for j := startNode + 1; j <= endNode; j++ {
				if !store[j] {
					val = false
					break
				}
			}
		}
		res[i] = val
	}

	return res
}
