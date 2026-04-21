// 1855. Maximum Distance Between a Pair of Values

package problems

import "fmt"

func Problem_1855() {
	nums1 := []int{55, 30, 5, 4, 2}
	nums2 := []int{100, 20, 10, 10, 5}
	fmt.Println(maxDistance(nums1, nums2))
}

func maxDistance(nums1 []int, nums2 []int) int {
	len1 := len(nums1)
	len2 := len(nums2)
	distance := 0

	for i, j := 0, 0; i < len1 && j < len2; {
		if nums1[i] <= nums2[j] {
			distance = max(distance, j-i)
			j++
		} else {
			i++
			if i > j {
				j = i
			}
		}
	}

	return distance
}
