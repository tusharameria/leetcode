package problems

import (
	"fmt"
	"sort"
)

func Problem_757() {
	intervals := [][]int{
		{1, 3},
		{1, 4},
		{2, 5},
		{3, 5},
	}
	fmt.Println(intersectionSizeTwo(intervals))
}

func intersectionSizeTwo(intervals [][]int) int {
	for len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] != intervals[j][1] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] > intervals[j][0]
	})

	ans := 2

	rightMost := intervals[0][1]
	leftMost := intervals[0][1] - 1

	for i := 1; i <= len(intervals)-1; i++ {
		right := intervals[i][1]
		left := intervals[i][0]
		if left > rightMost {
			ans += 2
			rightMost = right
			leftMost = right - 1
		} else if left > leftMost {
			ans += 1
			leftMost = rightMost
			rightMost = right
		}
	}
	return ans
}
