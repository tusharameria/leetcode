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

func BinarySearchInt(nums []int, val int) int {
	// this assumes nums is sorted in ascending order
	n := len(nums)
	left := 0
	right := n - 1
	for left <= right {
		index := left + (right-left+1)/2
		if nums[index] == val {
			return index
		} else if nums[index] > val {
			right = index - 1
		} else {
			left = index + 1
		}
	}
	return -1
}
