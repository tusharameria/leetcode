package topksortedstructures

import (
	"container/heap"
	"fmt"
	"sort"

	ds "github.com/tusharameria/leetcode/src/dataStructures"
)

// Smallest K pair sum
func KTopPairSum() {
	nums1 := []int{9, 5, 7}
	nums2 := []int{8, 4, 6}
	k := 3
	fmt.Println(KTopPairSumBrute(nums1, nums2, k))
}

func KTopPairSumBrute(nums1, nums2 []int, k int) []int {
	n1 := len(nums1)
	n2 := len(nums2)
	pairSum := make([]int, n1*n2)

	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			pairSum[(n1*i)+j] = nums1[i] + nums2[j]
		}
	}
	sort.Ints(pairSum)
	return pairSum[:k]
}

func KTopPairSumEff(nums1, nums2 []int, k int) []int {
	pairSum := make([]int, k)

	sort.Ints(nums1)
	sort.Ints(nums2)

	store := ds.MinIntHeapWithIndices{{nums1[0] + nums2[0], 0, 0}}
	heap.Init(&store)
	pairSumIdx := 0

	for {
		popped := heap.Pop(&store).([3]int)
		pairSum[pairSumIdx] = popped[0]
		if pairSumIdx+1 == k {
			break
		}
		pairSumIdx++
		nums1Idx := popped[1]
		nums2Idx := popped[2]
		heap.Push(&store, [3]int{nums1[nums1Idx+1] + nums1[nums2Idx], nums1Idx + 1, nums2Idx})
		heap.Push(&store, [3]int{nums1[nums1Idx] + nums1[nums2Idx+1], nums1Idx, nums2Idx + 1})
	}

	return pairSum
}
