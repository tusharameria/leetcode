// 3691. Maximum Total Subarray Value II

package problems

import (
	"container/heap"
	"fmt"

	ds "github.com/tusharameria/leetcode/src/dataStructures"
	segmenttrees "github.com/tusharameria/leetcode/src/dataStructures/segment_trees"
)

func Problem_3691() {
	nums := []int{4, 2, 5, 1}
	k := 4
	fmt.Println(maxTotalValue(nums, k))
}

func maxTotalValue(nums []int, k int) int64 {
	n := len(nums)
	ans := int64(0)
	rangedMinSegTree := segmenttrees.GenerateRangedMinSegmentTree(nums)
	rangedMaxSegTree := segmenttrees.GenerateRangedMaxSegmentTree(nums)

	store := ds.MaxIntHeapWithIndices{}
	heap.Init(&store)

	for i := 0; i < n-1; i++ {
		maxVal := segmenttrees.GetRangedMaxQueryValues(rangedMaxSegTree, [][]int{{i, n - 1}})[0]
		minVal := segmenttrees.GetRangedMinQueryValues(rangedMinSegTree, [][]int{{i, n - 1}})[0]
		heap.Push(&store, [3]int{maxVal - minVal, i, n - 1})
	}

	count := 0
	for store.Len() > 0 {
		popped := heap.Pop(&store).([3]int)
		left := popped[1]
		right := popped[2]
		ans += int64(popped[0])
		if count+1 == k {
			break
		}
		count++
		if left < right-1 {
			maxVal := segmenttrees.GetRangedMaxQueryValues(rangedMaxSegTree, [][]int{{left, right - 1}})[0]
			minVal := segmenttrees.GetRangedMinQueryValues(rangedMinSegTree, [][]int{{left, right - 1}})[0]
			heap.Push(&store, [3]int{maxVal - minVal, left, right - 1})
		}
	}

	return ans
}
