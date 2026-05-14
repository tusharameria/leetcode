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

// This assumes nums is sorted in ascending order
// Case I : If val is present in nums, then returns the indices of the first and last occurrence of val in nums
// Case II : If val is not present in nums, then returns the indices of the first element smaller than val and the first element greater than val in nums
// Case III : If val is smaller than all elements in nums, then returns -1 and 0
// Case IV : If val is greater than all elements in nums, then returns n-1 and n
func BinarySearchIntRange(nums []int, val int) (int, int) {
	n := len(nums)
	left := 0
	right := n - 1
	if val < nums[left] {
		return -1, 0
	}
	if val > nums[right] {
		return n - 1, n
	}
	for left < right {
		mid := left + (right-left+1)/2
		if nums[mid] == val {
			return mid - 1, mid + 1
		} else if nums[mid] > val {
			right = mid
		} else {
			left = mid
		}
		if left == right-1 {
			if nums[left] == val {
				return left - 1, left + 1
			} else if nums[right] == val {
				return right - 1, right + 1
			} else {
				return left, right
			}
		}
	}
	return left - 1, left + 1
}

func BinarySearchIntRangeLeft(nums []int, val int) int {
	n := len(nums)
	left := 0
	right := n - 1
	if val <= nums[left] {
		return -1
	}
	if val > nums[right] {
		return n - 1
	}
	for left < right {
		diff := right - left
		if diff == 1 {
			return left
		} else {
			mid := left + (diff+1)/2
			currentVal := nums[mid]
			if currentVal >= val {
				right = mid
			} else {
				left = mid
			}
		}
	}
	return -1
}

func GenerateBitArrayFromInt(n int) []int {
	if n == 0 {
		return []int{0}
	}
	bitArray := []int{}
	for n > 0 {
		bitArray = append(bitArray, n%2)
		n /= 2
	}
	return bitArray
}

func PowInt(base, exp int) int {
	result := 1
	for exp > 0 {
		if exp%2 == 1 {
			result *= base
		}
		base *= base
		exp /= 2
	}
	return result
}
