// 3737. Count Subarrays With Majority Element I

package problems

import "fmt"

func Problem_3737() {
	nums := []int{1, 2, 2, 3}
	target := 2
	fmt.Println(countMajoritySubarrays(nums, target))
}

func countMajoritySubarrays(nums []int, target int) int {
	n := len(nums)
	prefixCount := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		prefixCount[i] = prefixCount[i-1]
		if nums[i-1] == target {
			prefixCount[i]++
		}
	}

	// fmt.Println(prefixCount)

	res := 0

	for left := 0; left < n; left++ {
		for right := left; right < n; right++ {
			targetCount := prefixCount[right+1] - prefixCount[left]
			elementCount := right - left + 1
			if targetCount*2 > elementCount {
				res++
			}
		}
	}

	return res
}
