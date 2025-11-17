package problems

func Func_1() {
	list := GenerateLinkedList([]int{2, 4, 5, 4, 6})
	PrintLinkedList(list)
	Reverse(list)
}

func Reverse(list *ListNode) {
	if list != nil {
		end := list
		nextEnd := end.next
		for {
			if nextEnd == nil {
				break
			}
			end.next = nextEnd.next
			nextEnd.next = list
			list = nextEnd
			nextEnd = end.next
		}
	}
	PrintLinkedList(list)
}
