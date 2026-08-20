// 3069. Distribute Elements Into Two Arrays I

package problems

import "fmt"

func Problem_3069() {
	nums := []int{5, 4, 3, 8}
	fmt.Println(resultArray(nums))
}

// Constraints:

// --> 3 <= n <= 50
// --> 1 <= nums[i] <= 100
// --> All elements in nums are distinct.

func resultArray(nums []int) []int {
	n := len(nums)
	arr1 := make([]int, n)
	arr1[0] = nums[0]
	arr2 := make([]int, n)
	arr2[0] = nums[1]
	idx1, idx2 := 1, 1

	for i := 2; i < n; i++ {
		val := nums[i]
		if arr1[idx1-1] > arr2[idx2-1] {
			arr1[idx1] = val
			idx1++
		} else {
			arr2[idx2] = val
			idx2++
		}
	}

	return append(arr1[:idx1], arr2[:idx2]...)
}
