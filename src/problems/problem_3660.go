// 3660. Jump Game IX

package problems

import "fmt"

func Problem_3660() {
	nums := []int{9, 30, 16, 6, 36, 9}
	fmt.Println(maxValue(nums))
}

func maxValue(nums []int) []int {
	numLen := len(nums)
	res := make([]int, numLen)
	minArray := make([]int, numLen)
	maxArray := make([]int, numLen)
	minArray[numLen-1] = nums[numLen-1]
	maxArray[0] = nums[0]
	for i := 1; i < numLen; i++ {
		maxArray[i] = max(maxArray[i-1], nums[i])
		minArray[numLen-1-i] = min(minArray[numLen-i], nums[numLen-1-i])
	}
	for i := range nums {
		currentVal := nums[i]
		if currentVal == minArray[i] && currentVal == maxArray[i] {
			res[i] = maxArray[i]
		} else {
			tempMaxVal := maxArray[i]
			nextMaxVal := maxArray[BinarySearchIntRangeLeft(minArray, tempMaxVal)]
			for nextMaxVal > tempMaxVal {
				newIndex := BinarySearchIntRangeLeft(maxArray, nextMaxVal) + 1
				tempMaxVal = maxArray[newIndex]
				nextMaxVal = maxArray[BinarySearchIntRangeLeft(minArray, tempMaxVal)]
			}
			res[i] = nextMaxVal
		}
	}
	return res
}
