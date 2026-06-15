// 2095. Delete the Middle Node of a Linked List

package problems

func Problem_2095() {}

func deleteMiddle(head *ListNode) *ListNode {
	lead := head.Next
	if lead == nil {
		return nil
	}

	trail := head
	lead = lead.Next

	for lead != nil && lead.Next != nil {
		lead = lead.Next.Next
		trail = trail.Next
	}

	trail.Next = trail.Next.Next

	return head
}
