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
	maxMap := make(map[int]int)
	minArray[numLen-1] = nums[numLen-1]
	maxArray[0] = nums[0]
	for i := 1; i < numLen; i++ {
		maxArray[i] = max(maxArray[i-1], nums[i])
		minArray[numLen-1-i] = min(minArray[numLen-i], nums[numLen-1-i])
	}
	for i := range nums {
		currentVal := nums[i]
		minVal := minArray[i]
		maxVal := maxArray[i]
		if currentVal == minVal && currentVal == maxVal {
			res[i] = maxVal
		} else {
			if maxMap[maxVal] > 0 {
				res[i] = maxMap[maxVal]
				continue
			}
			tempArr := []int{maxVal}
			tempMaxVal := maxVal
			nextMaxVal := maxArray[BinarySearchIntRangeLeft(minArray, tempMaxVal)]
			if tempMaxVal != nextMaxVal {
				tempArr = append(tempArr, nextMaxVal)
			}
			for nextMaxVal > tempMaxVal {
				newIndex := BinarySearchIntRangeLeft(maxArray, nextMaxVal) + 1
				tempMaxVal = maxArray[newIndex]
				nextMaxVal = maxArray[BinarySearchIntRangeLeft(minArray, tempMaxVal)]
				tempArr = append(tempArr, nextMaxVal)
			}
			res[i] = nextMaxVal
			for _, val := range tempArr {
				maxMap[val] = nextMaxVal
			}
		}
	}
	return res
}
