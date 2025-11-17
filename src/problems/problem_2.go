package problems

func Problem_2() {
	l1 := []int{2, 4, 3}
	l2 := []int{5, 6, 4}
	PrintLinkedList(addTwoNumbers(GenerateLinkedList(l1), GenerateLinkedList(l2)))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	start := res
	up := l1
	down := l2
	carry := 0
	for {
		if up == nil && down == nil && carry == 0 {
			break
		} else {
			val1 := 0
			val2 := 0
			if up != nil {
				val1 = up.val
				up = up.next
			}
			if down != nil {
				val2 = down.val
				down = down.next
			}
			sum := val1 + val2 + carry
			carry = int(sum / 10)

			start.next = &ListNode{
				val: sum % 10,
			}
			start = start.next

		}
	}
	return res.next
}
