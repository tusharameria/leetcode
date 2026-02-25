package problems

import (
	"fmt"
	"math/bits"
	"sort"
)

// 1356. Sort Integers by The Number of 1 Bits

func Problem_1356() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("res : ", sortByBits(arr))
}

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		iCount := bits.OnesCount(uint(arr[i]))
		jCount := bits.OnesCount(uint(arr[j]))
		if iCount == jCount {
			return arr[i] < arr[j]
		}
		return iCount < jCount
	})
	return arr
}
