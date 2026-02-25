package problems

import (
	"cmp"
	"fmt"
	"math/bits"
	"slices"
)

// 1356. Sort Integers by The Number of 1 Bits

func Problem_1356() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("res : ", sortByBits(arr))
}

func sortByBits(arr []int) []int {
	slices.SortFunc(arr, func(a, b int) int {
		iCount := bits.OnesCount(uint(a))
		jCount := bits.OnesCount(uint(b))
		if iCount == jCount {
			return cmp.Compare(a, b)
		}
		return cmp.Compare(iCount, jCount)
	})
	return arr
}
