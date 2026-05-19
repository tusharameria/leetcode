// 2540. Minimum Common Value

package problems

import "fmt"

func Problem_2540() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4}
	fmt.Println(getCommon(nums1, nums2))
}

func getCommon(nums1 []int, nums2 []int) int {
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] {
			return nums1[i]
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return -1
}
