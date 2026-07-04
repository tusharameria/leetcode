// 2492. Minimum Score of a Path Between Two Cities

package problems

import (
	"fmt"
	"math"
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
	visited := make([]bool, n+1)
	adjLists := make([][][2]int, n+1)

	for _, node := range roads {
		node1, node2, cost := node[0], node[1], node[2]
		adjLists[node1] = append(adjLists[node1], [2]int{node2, cost})
		adjLists[node2] = append(adjLists[node2], [2]int{node1, cost})
	}

	stack := adjLists[1]
	visited[1] = true
	res := math.MaxInt
	for len(stack) > 0 {
		adjList := stack[0]
		stack = stack[1:]
		newSource := adjList[0]
		res = min(res, adjList[1])
		for _, ls := range adjLists[newSource] {
			if !visited[ls[0]] {
				stack = append(stack, [2]int{ls[0], ls[1]})
			}
		}
		visited[newSource] = true
	}
	return res
}
