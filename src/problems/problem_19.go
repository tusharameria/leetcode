package problems

func Problem_19() {
	head := GenerateLinkedList([]int{1, 2, 3, 4, 5})
	n := 5
	PrintLinkedList((removeNthFromEnd(head, n)))
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	trail := head
	lead := head
	i := 0
	for {
		if lead.Next == nil {
			if i < n {
				return head.Next
			}
			trail.Next = trail.Next.Next
			break
		}
		lead = lead.Next
		if n <= i {
			trail = trail.Next
		}
		i++
	}
	return head
}
