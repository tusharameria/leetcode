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
	arr1 := make([]uint8, n)
	arr1[0] = uint8(nums[0])
	arr2 := make([]uint8, n)
	arr2[0] = uint8(nums[1])
	idx1, idx2 := 1, 1

	for i := 2; i < n; i++ {
		val := uint8(nums[i])
		if arr1[idx1-1] > arr2[idx2-1] {
			arr1[idx1] = val
			idx1++
		} else {
			arr2[idx2] = val
			idx2++
		}
	}

	for i := 0; i < idx1; i++ {
		nums[i] = int(arr1[i])
	}

	for i := 0; i < idx2; i++ {
		nums[i+idx1] = int(arr2[i])
	}

	return nums
}
