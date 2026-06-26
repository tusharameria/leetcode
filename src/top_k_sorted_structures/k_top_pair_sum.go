package topksortedstructures

import (
	"container/heap"
	"fmt"
	"sort"

	ds "github.com/tusharameria/leetcode/src/dataStructures"
)

// Smallest K pair sum
func KTopPairSum() {
	nums1 := []int{1, 2, 4, 5, 6}
	nums2 := []int{3, 5, 7, 9}
	k := 20
	fmt.Println(KTopPairSumEff(nums1, nums2, k))
}

func KTopPairSumBrute(nums1, nums2 []int, k int) [][]int {
	n1 := len(nums1)
	n2 := len(nums2)
	pairSum := make([][]int, n1*n2)

	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			pairSum[(n1*i)+j] = []int{nums1[i] + nums2[j], i, j}
		}
	}
	sort.Slice(pairSum, func(i, j int) bool {
		return pairSum[i][0] < pairSum[j][0]
	})

	pairSum = pairSum[:k]

	for i := 0; i < k; i++ {
		pairSum[i] = pairSum[i][:1]
	}
	return pairSum
}

func KTopPairSumEff(nums1, nums2 []int, k int) [][]int {
	n1 := len(nums1)
	n2 := len(nums2)
	pairSum := make([][]int, k)

	// If not sorted
	// sort.Ints(nums1)
	// sort.Ints(nums2)

	visited := make(map[[2]int]bool)

	store := ds.MinIntHeapWithIndices{{nums1[0] + nums2[0], 0, 0}}
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
