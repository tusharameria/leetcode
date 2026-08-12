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
	maxLen := 1
	left := 0
	store := make(map[int]int)
	store[nums[left]]++
	right := 1
	for ; right < len(nums); right++ {
		currentVal := nums[right]
		store[currentVal]++
		if store[currentVal] > k {
			maxLen = max(maxLen, right-left)
			fmt.Println("left : ", left)
			fmt.Println("right : ", right)
			fmt.Println("currentVal : ", currentVal)
			for j := left; j < right; j++ {
				leftVal := nums[j]
				store[leftVal]--
				if leftVal == currentVal {
					left = j + 1
					break
				}
			}
		}
	}
	fmt.Println(maxLen)
	fmt.Println(left)
	fmt.Println(right)
	maxLen = max(maxLen, right-left)

	return maxLen
}
