package graph

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

func Dijkstra() {
	edges := [][]int{
		{2, 1, 1},
		{1, 3, 2},
		{0, 2, 5},
		{2, 3, 6},
		{0, 1, 7},
	}
	online := []bool{true, true, true, true}
	k := int64(11)
	minPathValue := 2

	// Main Algo
	n := len(online)
	edgesLen := len(edges)

	adjacentList := make([][][2]int, n)

	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	lastNodeTarget := false
	for i := 0; i < edgesLen; i++ {
		edge := edges[i]
		if online[edge[0]] && online[edge[1]] {
			adjacentList[edge[0]] = append(adjacentList[edge[0]], [2]int{edge[1], edge[2]})
			if edge[1] == n-1 {
				lastNodeTarget = true
			}
		}
	}
	if !lastNodeTarget {
		fmt.Println(false)
	}
	fmt.Println(dijkstraPossible(adjacentList, minPathValue, k))
}

func dijkstraPossible(adjacentList [][][2]int, minPathValue int, totalCost int64) bool {
	n := len(adjacentList)
	scores := make([]int, n)
	inf := math.MaxInt
	for i := range n {
		scores[i] = inf
	}
	scores[0] = 0

	h := &MinIntHeapDijk{}
	heap.Init(h)
	heap.Push(h, [2]int{0, 0})

	for h.Len() > 0 {
		currentNode := h.Pop().([2]int)
		adArray := adjacentList[currentNode[0]]
		for i := 0; i < len(adArray); i++ {
			targetNode, cost := adArray[i][0], adArray[i][1]
			targetCost := scores[currentNode[0]] + cost
			if cost >= minPathValue && targetCost <= int(totalCost) && targetCost < scores[targetNode] {
				if targetNode == n-1 {
					return true
				}
				scores[targetNode] = targetCost
				heap.Push(h, [2]int{targetNode, targetCost})
			}
		}
	}

	return false
}

type MinIntHeapDijk [][2]int

func (h MinIntHeapDijk) Len() int           { return len(h) }
func (h MinIntHeapDijk) Less(i, j int) bool { return h[i][1] < h[j][1] } // min-heap
func (h MinIntHeapDijk) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinIntHeapDijk) Push(x any) {
	*h = append(*h, x.([2]int))
}
func (h *MinIntHeapDijk) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
