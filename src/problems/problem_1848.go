// 1848. Minimum Distance to the Target Element

package problems

import "fmt"

func Problem_1848() {
	nums := []int{}
	target := 0
	start := 0
	fmt.Println(getMinDistance(nums, target, start))
}

func getMinDistance(nums []int, target int, start int) int {
	closest := -1
	for i := 0; i <= len(nums)-1; i++ {
		if closest != -1 && closest < start && i == start+(start-closest) {
			return start - closest
		}
		if nums[i] == target {
			if i >= start {
				return i - start
			} else {
				closest = i
			}
		}
	}
	return start - closest
}
