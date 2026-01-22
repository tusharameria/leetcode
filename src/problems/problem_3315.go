package problems

import (
	"fmt"
	"math"
)

// 3315. Construct the Minimum Bitwise Array II

func Problem_3315() {
	nums := []int{2, 3, 5, 7, 11, 13, 17, 19, 23}
	fmt.Printf("%v\n", minBitwiseArray(nums))
}

func minBitwiseArray(nums []int) []int {
	for i := 0; i <= len(nums)-1; i++ {
		nums[i] = getMinNumber(nums[i])
	}
	return nums
}

// Returns Minimum Number a so that bitwise a OR a+1 is equal to n
func getMinNumber(n int) int {
	if n == 2 {
		return -1
	}
	ans := 0
	bitArray := GenerateBitArrayFromInt(n)
	zeroFound := false
	for i := 0; i <= len(bitArray)-1; i++ {
		if zeroFound {
			ans += bitArray[i] * int(math.Pow(2, float64(i)))
			continue
		}
		if i+1 <= len(bitArray)-1 && bitArray[i+1] == 1 {
			ans += bitArray[i] * int(math.Pow(2, float64(i)))
		} else {
			zeroFound = true
		}
	}
	return ans
}
