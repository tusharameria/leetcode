// 3689. Maximum Total Subarray Value I

package problems

func Problem_3689() {}

func maxTotalValue(nums []int, k int) int64 {
	maxVal := nums[0]
	minVal := maxVal

	for i := 1; i < len(nums); i++ {
		val := nums[i]
		maxVal = max(maxVal, val)
		minVal = min(minVal, val)
	}

	return (int64((maxVal - minVal) * k))
}
