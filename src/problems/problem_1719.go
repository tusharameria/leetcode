// 1979. Find Greatest Common Divisor of Array

package problems

import (
	"fmt"
	"math"
)

func Problem_1719() {
	nums := []int{2, 5, 6, 9, 10}
	fmt.Println(findGCD(nums))
}

func findGCD(nums []int) int {
	var minVal int = math.MaxUint16
	var maxVal int = 1
	for _, val := range nums {
		minVal = min(minVal, val)
		maxVal = max(maxVal, val)
	}
	for maxVal%minVal != 0 {
		minVal, maxVal = maxVal%minVal, minVal
	}
	return minVal
}
