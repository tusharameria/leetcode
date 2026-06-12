// 3558. Number of Ways to Assign Edge Weights I

package problems

import (
	"fmt"
)

func Problem_3558() {
	edges := [][]int{{1, 2}, {1, 3}, {3, 4}, {3, 5}}
	fmt.Println(edges)
}

func assignEdgeWeights(edges [][]int) int {
	n := len(edges) + 1
	graph := make([][]int, n+1)

	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}

	var dfs func(node, prev int) int

	dfs = func(node, prev int) int {
		dist := 0
		for _, val := range graph[node] {
			if val != prev {
				dist = max(dist, dfs(val, node)+1)
			}
		}
		return dist
	}

	return pow(2, dfs(1, 0)-1)
}

func pow(base, exp int) int {
	const mod = 1_000_000_007
	res := 1

	for exp > 0 {
		if exp%2 > 0 {
			res = res * base % mod
		}
		base = base * base % mod
		exp /= 2
	}

	return res
}
