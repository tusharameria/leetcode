package problems

import (
	"fmt"
)

// 1200. Minimum Absolute Difference

func Problem_1200() {
	// arr := []int{4, 2, 1, 3}
	// arr := []int{1, 3, 6, 10, 15}
	arr := []int{3, 8, -10, 23, 19, -4, -14, 27}
	fmt.Println("res : ", minimumAbsDifference(arr))
}

func minimumAbsDifference(arr []int) [][]int {
	n := len(arr)

	// Find min and max values in the original array
	minVal := arr[0]
	maxVal := arr[0]
	for i := 0; i <= n-1; i++ {
		minVal = min(minVal, arr[i])
		maxVal = max(maxVal, arr[i])
	}

	// Length of dp array
	m := maxVal - minVal + 1
	dp := make([]bool, m)

	// Initialise dp array
	for i := 0; i <= n-1; i++ {
		dp[arr[i]-minVal] = true

	}

	res := make([][]int, 0, n)

	l := 0
	minDiff := m - 1

	for r := 1; r <= m-1; r++ {
		if dp[r] {
			if r-l <= minDiff {
				if r-l < minDiff {
					minDiff = r - l
					res = res[:0] // Clear the result array
				}
				res = append(res, []int{l + minVal, r + minVal})
			}
			l = r
		}
	}
	return res
}
