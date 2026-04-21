// 3488. Closest Equal Element Queries

package problems

import "fmt"

func Problem_3488() {
	nums := []int{1, 3, 1, 4, 1, 3, 2}
	queries := []int{0, 3, 5}
	fmt.Println(solveQueries(nums, queries))
}

func solveQueries(nums []int, queries []int) []int {
	res := make([]int, len(queries))
	for i := 0; i <= len(queries)-1; i++ {
		res[i] = findClosest(nums, queries[i])
	}
	return res
}

func findClosest(nums []int, startIndex int) int {
	distance := -1
	for i := 0; i <= len(nums)-1; i++ {
		if i != startIndex {
			if nums[i] == nums[startIndex] {
				if distance == -1 {
					distance = min(abs(i-startIndex), len(nums)-abs(i-startIndex))
				} else {
					distance = min(distance, min(abs(i-startIndex), len(nums)-abs(i-startIndex)))
				}
				if distance == 1 {
					return 1
				}
			}
		}
	}
	return distance
}
