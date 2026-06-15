// 2130. Maximum Twin Sum of a Linked List

package problems

import "fmt"

func Problem_2130() {
	head := []int{5, 4, 2, 2}
	fmt.Println(pairSum(GenerateLinkedList(head)))
}

func pairSum(head *ListNode) int {
	trail := head
	lead := head.Next
	for lead.Next != nil {
		trail = trail.Next
		lead = lead.Next.Next
	}
	lead = trail.Next
	forward := trail.Next
	for forward.Next != nil {
		buff := forward.Next
		forward.Next = buff.Next
		buff.Next = lead
		lead = buff
	}

	maxVal := 0
	for lead != nil {
		maxVal = max(maxVal, head.Val+lead.Val)
		head = head.Next
		lead = lead.Next
	}
	return maxVal
}
