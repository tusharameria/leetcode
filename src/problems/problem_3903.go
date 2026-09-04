// 3903. Smallest Stable Index I

package problems

func Problem_3903() {}

// Constraints:

// --> 1 <= nums.length <= 100
// --> 0 <= nums[i] <= 10^9
// --> 0 <= k <= 10^9

func firstStableIndex(nums []int, k int) int {
	n := len(nums)
	resIdx := n
	maxArr := make([]int, n)
	maxArr[0] = nums[0]
	minVal := nums[n-1]
	for i := 1; i < n; i++ {
		maxArr[i] = max(maxArr[i-1], nums[i])
	}
	if maxArr[n-1]-minVal <= k {
		resIdx = n - 1
	}
	for i := n - 2; i >= 0; i-- {
		minVal = min(minVal, nums[i])
		if maxArr[i]-minVal <= k {
			resIdx = i
		}
	}
	if resIdx == n {
		return -1
	}
	return resIdx
}
