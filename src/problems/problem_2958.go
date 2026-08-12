// 2958. Length of Longest Subarray With at Most K Frequency

package problems

import "fmt"

func Problem_2958() {
	nums := []int{3, 1, 1}
	k := 1
	fmt.Println(maxSubarrayLength(nums, k))
}

// Constraints :

// ---> 1 <= nums.length <= 10^5
// ---> 1 <= nums[i] <= 10^9
// ---> 1 <= k <= nums.length

func maxSubarrayLength(nums []int, k int) int {
	newK := uint32(k)
	n := uint32(len(nums))
	windFreqs := make(map[int]uint32, n)
	var left uint32 = 0
	var res uint32 = 0
	var right uint32 = 0

	for ; right < n; right++ {
		windFreqs[nums[right]]++

		for windFreqs[nums[right]] > newK {
			windFreqs[nums[left]]--
			left++
		}

		res = max(res, right-left+1)
	}

	return int(res)
}
