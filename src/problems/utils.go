package problems

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func GenerateLinkedList(arr []int) *ListNode {
	var nextList *ListNode
	for i := len(arr) - 1; i >= 0; i-- {
		buff := &ListNode{
			val:  arr[i],
			next: nextList,
		}
		nextList = buff
	}
	return nextList
}

func PrintLinkedList(list *ListNode) {
	for {
		if list != nil {
			fmt.Printf("%d ", list.val)
			list = list.next
		} else {
			fmt.Println()
			break
		}
	}
}
