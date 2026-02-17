package problems

import (
	"fmt"
	"sort"
)

// 1200. Minimum Absolute Difference

func Problem_1200() {
	arr := []int{4, 2, 1, 3}
	// arr := []int{1, 3, 6, 10, 15}
	// arr := []int{3, 8, -10, 23, 19, -4, -14, 27}
	fmt.Println("res : ", minimumAbsDifference(arr))
}

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	res := make([][]int, 0)
	min := arr[1] - arr[0]
	res = append(res, []int{arr[0], arr[1]})
	for i := 2; i <= len(arr)-1; i++ {
		diff := arr[i] - arr[i-1]
		if diff == min {
			res = append(res, []int{arr[i-1], arr[i]})
		} else if diff < min {
			min = diff
			res = make([][]int, 0)
			res = append(res, []int{arr[i-1], arr[i]})
		}
	}
	return res
}
