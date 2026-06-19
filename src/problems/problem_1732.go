// 1732. Find the Highest Altitude

package problems

import "fmt"

func Problem_1732() {
	gain := []int{-5, 1, 5, 0, -7}
	fmt.Println(largestAltitude(gain))
}

func largestAltitude(gain []int) int {
	maxVal := 0
	currSum := 0
	for _, val := range gain {
		currSum += val
		maxVal = max(maxVal, currSum)
	}
	return maxVal
}
