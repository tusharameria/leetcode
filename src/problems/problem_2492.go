// 2492. Minimum Score of a Path Between Two Cities

package problems

import (
	"fmt"
)

func Problem_2492() {
	n := 4
	roads := [][]int{
		{1, 2, 9},
		{2, 3, 6},
		{2, 4, 5},
		{1, 4, 7},
	}
	fmt.Println(minScore(n, roads))
}

func minScore(n int, roads [][]int) int {
	g := make([][][2]int, n+1)

	for _, node := range roads {
		node1, node2, cost := node[0], node[1], node[2]
		g[node1] = append(g[node1], [2]int{node2, cost})
		g[node2] = append(g[node2], [2]int{node1, cost})
	}

	visited := make([]bool, n+1)
	stack := []int{1}
	visited[1] = true

	res := 100000000
	for i := 0; i < len(stack); i++ {
		newSource := stack[i]
		for _, ls := range g[newSource] {
			if !visited[ls[0]] {
				stack = append(stack, ls[0])
				visited[newSource] = true
				res = min(res, ls[1])
			}
		}
	}
	return res
}
