// 3558. Number of Ways to Assign Edge Weights I

package problems

import (
	"fmt"
	"sort"
)

func Problem_3558() {
	edges := [][]int{{1, 2}, {1, 3}, {3, 4}, {3, 5}}
	fmt.Println(edges)
}

const MOD = 1_000_000_007

var depth [100_001]uint32

func assignEdgeWeightsOld(edges [][]int) int {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][0] < edges[j][0]
	})
	maxDepth := uint32(0)
	for _, e := range edges {
		parent := min(e[0], e[1])
		child := max(e[0], e[1])
		depth[child] = depth[parent] + 1
		maxDepth = max(maxDepth, depth[child])
	}
	return pow2(maxDepth - 1)
}

func pow2(n uint32) int {
	result := 1
	base := 2
	for n > 0 {
		if n&1 == 1 {
			result = result * base % MOD
		}
		base = base * base % MOD
		n >>= 1
	}
	return result
}
