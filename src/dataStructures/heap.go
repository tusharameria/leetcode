package ds

import (
	"container/heap"
	"fmt"
)

type MaxIntHeap []int
type MinIntHeap []int

func (h MaxIntHeap) Len() int           { return len(h) }
func (h MaxIntHeap) Less(i, j int) bool { return h[i] > h[j] } // max-heap
func (h MaxIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h MinIntHeap) Len() int           { return len(h) }
func (h MinIntHeap) Less(i, j int) bool { return h[i] < h[j] } // min-heap
func (h MinIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxIntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinIntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxIntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *MinIntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	h := &MaxIntHeap{2, 1, 5}
	heap.Init(h)

	heap.Push(h, 3)
	fmt.Println(heap.Pop(h)) // 1
	fmt.Println(heap.Pop(h)) // 2
}
