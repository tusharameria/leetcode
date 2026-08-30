// 2091. Removing Minimum and Maximum From Array

package problems

import "fmt"

func Problem_2091() {
	nums := []int{0, -4, 19, 1, 8, -2, -3, 5}
	fmt.Println(minimumDeletions(nums))
}

// Constraints:

// --> 1 <= nums.length <= 10^5
// --> -10^5 <= nums[i] <= 10^5
// --> The integers in nums are distinct.

func minimumDeletions(nums []int) int {
	n := len(nums)
	maxNum := -1_00_000
	minNum := 1_00_000
	maxIdx := 0
	minIdx := 0
	for i := 0; i < n; i++ {
		num := nums[i]
		if maxNum < num {
			maxIdx = i
			maxNum = num
		}
		if minNum > num {
			minIdx = i
			minNum = num
		}
	}

	leftIdx := maxIdx
	rightIdx := minIdx

	if maxIdx > minIdx {
		leftIdx, rightIdx = minIdx, maxIdx
	}

	return min(rightIdx+1, n-leftIdx, leftIdx+1+n-rightIdx)
}
