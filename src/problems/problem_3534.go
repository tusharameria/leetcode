// 3534. Path Existence Queries in a Graph II

package problems

import (
	"fmt"
	"sort"
)

func Problem_3534() {
	n := 5
	nums := []int{5, 3, 1, 9, 10}
	maxDiff := 2
	queries := [][]int{{0, 1}, {0, 2}, {2, 3}, {4, 3}}
	fmt.Println(pathExistenceQueries(n, nums, maxDiff, queries))
}

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []int {
	order := make([]int, n)
	for i := range n {
		order[i] = i
	}
	sort.Slice(order, func(i, j int) bool {
		return nums[order[i]] < nums[order[j]]
	})
	vals := make([]int, n)
	pos := make([]int, n)
	for i := range n {
		pos[order[i]] = i
		vals[i] = nums[order[i]]
	}

	reach := make([]int, n)
	j := 0
	for i := 0; i < n; i++ {
		if j < i {
			j = i
		}
		for j+1 < n && vals[j+1]-vals[i] <= maxDiff {
			j++
		}
		reach[i] = j
	}

	fmt.Println(order)
	fmt.Println(vals)
	fmt.Println(pos)
	fmt.Println(reach)
	comp := make([]int, n)
	for i := 1; i < n; i++ {
		comp[i] = comp[i-1]
		if vals[i]-vals[i-1] > maxDiff {
			comp[i]++
		}
	}

	LOG := 1
	for (1 << LOG) < n {
		LOG++
	}
	jump := make([][]int, LOG)
	jump[0] = make([]int, n)
	copy(jump[0], reach)
	fmt.Println("=============")
	for k := 1; k < LOG; k++ {
		jump[k] = make([]int, n)
		for i := 0; i < n; i++ {
			jump[k][i] = jump[k-1][jump[k-1][i]]
		}
	}

	for _, val := range jump {
		fmt.Println(val)
	}

	ans := make([]int, len(queries))
	for qi, q := range queries {
		pu, pv := pos[q[0]], pos[q[1]]
		if comp[pu] != comp[pv] {
			ans[qi] = -1
			continue
		}
		if pu == pv {
			ans[qi] = 0
			continue
		}
		if pu > pv {
			pu, pv = pv, pu
		}
		dist, cur := 0, pu
		for k := LOG - 1; k >= 0; k-- {
			if jump[k][cur] < pv {
				cur = jump[k][cur]
				dist += 1 << k
			}
		}
		if cur < pv {
			dist++
		}
		ans[qi] = dist
	}
	return ans
}
