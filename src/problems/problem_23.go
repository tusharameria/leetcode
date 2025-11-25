package problems

func Problem_23() {
	lists := []*ListNode{
		GenerateLinkedList([]int{2, 3, 4, 5}),
	}
	PrintLinkedList(mergeKLists(lists))
}

func mergeKLists(lists []*ListNode) *ListNode {
	return nil
}
