// 373. Find K Pairs with Smallest Sums

package problems

import "container/heap"

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	n1 := len(nums1)
	n2 := len(nums2)
	pairSum := make([][]int, k)

	visited := make(map[[2]int]bool)

	store := MinIntHeapWithIndices{{nums1[0] + nums2[0], 0, 0}}
	heap.Init(&store)
	pairSumIdx := 0

	for {
		popped := heap.Pop(&store).([3]int)
		nums1Idx := popped[1]
		nums2Idx := popped[2]
		pairSum[pairSumIdx] = []int{nums1[nums1Idx], nums2[nums2Idx]}
		if pairSumIdx+1 == k {
			break
		}
		pairSumIdx++
		if nums1Idx+1 < n1 && !visited[[2]int{nums1Idx + 1, nums2Idx}] {
			heap.Push(&store, [3]int{nums1[nums1Idx+1] + nums2[nums2Idx], nums1Idx + 1, nums2Idx})
			visited[[2]int{nums1Idx + 1, nums2Idx}] = true
		}
		if nums2Idx+1 < n2 && !visited[[2]int{nums1Idx, nums2Idx + 1}] {
			heap.Push(&store, [3]int{nums1[nums1Idx] + nums2[nums2Idx+1], nums1Idx, nums2Idx + 1})
			visited[[2]int{nums1Idx, nums2Idx + 1}] = true
		}
	}

	return pairSum
}

type MinIntHeapWithIndices [][3]int

func (h MinIntHeapWithIndices) Len() int           { return len(h) }
func (h MinIntHeapWithIndices) Less(i, j int) bool { return h[i][0] < h[j][0] } // min-heap
func (h MinIntHeapWithIndices) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinIntHeapWithIndices) Push(x any) {
	*h = append(*h, x.([3]int))
}
func (h *MinIntHeapWithIndices) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
