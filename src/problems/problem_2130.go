// 2130. Maximum Twin Sum of a Linked List

package problems

func pairSum(head *ListNode) int {
	nums := make([]int, 1_00_000)
	idx := -1
	for head != nil {
		idx++
		nums[idx] = head.Val
		head = head.Next
	}
	maxVal := 0
	for i := 0; i <= idx/2; i++ {
		maxVal = max(maxVal, nums[i]+nums[idx-i])
	}
	return maxVal
}
