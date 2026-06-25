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
	pre := make([]uint16, (2*n)+1)
	pre[n] = 1

	cnt := n
	ans := 0
	preSum := 0

	for _, val := range nums {
		if val == target {
			preSum += int(pre[cnt])
			cnt++
			pre[cnt]++
		} else {
			cnt--
			preSum -= int(pre[cnt])
			pre[cnt]++
		}
		ans += preSum
	}

	return ans
}
