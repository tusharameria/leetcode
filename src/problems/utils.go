package problems

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func GenerateLinkedList(arr []int) *ListNode {
	var nextList *ListNode
	for i := len(arr) - 1; i >= 0; i-- {
		buff := &ListNode{
			Val:  arr[i],
			Next: nextList,
		}
		nextList = buff
	}
	return nextList
}

func PrintLinkedList(list *ListNode) {
	for {
		if list != nil {
			fmt.Printf("%d ", list.Val)
			list = list.Next
		} else {
			fmt.Println()
			break
		}
	}
}

func ReverseLinkedList(list *ListNode) {
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
