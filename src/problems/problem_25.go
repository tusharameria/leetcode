package problems

func Problem_25() {
	head := GenerateLinkedList([]int{1, 2, 3, 5, 6, 7, 8})
	k := 4
	PrintLinkedList(reverseKGroup(head, k))
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	// lead := head
	if k == 1 {
		return head
	}
	end := head
	nextEnd := end.Next
	for i := 0; i < k-1; i++ {
		end.Next = nextEnd.Next
		nextEnd.Next = head
		head = nextEnd
		nextEnd = end.Next
	}

	trail := end
	lead := nextEnd

	i := 0
	for lead != nil {
		if (i+1)%k != 0 {
			lead = lead.Next
		} else {
			end = trail.Next
			nextEnd = end.Next
			for j := 0; j < k-1; j++ {
				end.Next = nextEnd.Next
				nextEnd.Next = trail.Next
				trail.Next = nextEnd
				nextEnd = end.Next
			}
			trail = end
			lead = nextEnd
		}
		i++
	}
	return head
}
