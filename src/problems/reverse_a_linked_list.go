package problems

func Func_1() {
	list := GenerateLinkedList([]int{2, 4, 5, 4, 6})
	PrintLinkedList(list)
	Reverse(list)
}

func Reverse(list *ListNode) {
	if list != nil {
		end := list
		nextEnd := end.Next
		for {
			if nextEnd == nil {
				break
			}
			end.Next = nextEnd.Next
			nextEnd.Next = list
			list = nextEnd
			nextEnd = end.Next
		}
	}
	PrintLinkedList(list)
}
