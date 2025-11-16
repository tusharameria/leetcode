package problems

import "fmt"

type LinkedList struct {
	val  int
	next *LinkedList
}

func GenerateLinkedList(arr []int) *LinkedList {
	var nextList *LinkedList
	for i := len(arr) - 1; i >= 0; i-- {
		buff := &LinkedList{
			val:  arr[i],
			next: nextList,
		}
		nextList = buff
	}
	return nextList
}

func PrintLinkedList(list *LinkedList) {
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
