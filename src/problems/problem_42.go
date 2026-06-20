// 42. Trapping Rain Water

package problems

import "fmt"

func Problem_42() {
	height := []int{4, 2, 0, 3, 2, 5}
	fmt.Println(trap(height))
}

func trap(height []int) int {
	n := len(height)
	leftMax := make([]int, n)
	rightMax := make([]int, n)
	leftMaxVal := 0
	rightMaxVal := 0
	for i := 0; i < n; i++ {
		leftMaxVal = max(leftMaxVal, height[i])
		rightMaxVal = max(rightMaxVal, height[n-1-i])
		leftMax[i] = leftMaxVal
		rightMax[n-1-i] = rightMaxVal
	}
	trapped := 0

	for i := 0; i < n; i++ {
		trapped += max(0, min(leftMax[i], rightMax[i])-height[i])
	}

	return trapped
}
