package problems

func Problem_21() {
	list1 := GenerateLinkedList([]int{1, 2, 4})
	list2 := GenerateLinkedList([]int{1, 3, 4})

	PrintLinkedList(mergeTwoLists(list1, list2))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	final := &ListNode{}
	res := final
	up := list1
	down := list2

	for {
		if up == nil || down == nil {
			if up == nil {
				res.Next = down
			} else {
				res.Next = up
			}
			break
		} else {
			if up.Val < down.Val {
				res.Next = &ListNode{
					Val:  up.Val,
					Next: up.Next,
				}
				up = up.Next
			} else {
				res.Next = &ListNode{
					Val:  down.Val,
					Next: down.Next,
				}
				down = down.Next
			}
			res = res.Next
		}
	}
	return final.Next
}

func mergeTwoListsEff(list1 *ListNode, list2 *ListNode) *ListNode {
	ans := &ListNode{}
	head := ans
	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			head.Next = list2
			list2 = list2.Next
		} else {
			head.Next = list1
			list1 = list1.Next
		}
		head = head.Next
	}
	if list1 == nil {
		head.Next = list2
	}
	if list2 == nil {
		head.Next = list1
	}
	return ans.Next
}
