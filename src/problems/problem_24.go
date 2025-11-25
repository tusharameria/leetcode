package problems

func Problem_24() {
	head := GenerateLinkedList([]int{1, 2, 3, 4, 5})
	PrintLinkedList((swapPairs(head)))
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ans := head.Next
	head.Next = head.Next.Next
	ans.Next = head
	for head.Next != nil && head.Next.Next != nil {
		buff := head.Next.Next
		head.Next.Next = head.Next.Next.Next
		buff.Next = head.Next
		head.Next = buff
		head = head.Next.Next
	}
	return ans
}

// 2 -> 1 -> 3 -> 4 -> 5
// 2 -> 1
