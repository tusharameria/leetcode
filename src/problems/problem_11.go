package problems

import (
	"fmt"
)

func Problem_11() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	maxArea := 0
	n := len(height)
	leftIndex, rightIndex := 0, n-1
	for leftIndex < rightIndex {
		maxArea = max(maxArea, (rightIndex-leftIndex)*min(height[leftIndex], height[rightIndex]))
		if height[leftIndex] < height[rightIndex] {
			leftIndex++
		} else {
			rightIndex--
		}
	}
	return maxArea
}
