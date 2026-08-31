// 2058. Find the Minimum and Maximum Number of Nodes Between Critical Points

package problems

import "fmt"

func Problem_2058() {
	arr := []int{7, 7, 10, 1, 7, 1, 2, 1, 5}
	head := GenerateLinkedList(arr)
	fmt.Println(nodesBetweenCriticalPoints(head))
}

// Constraints:

// --> The number of nodes in the list is in the range [2, 10^5].
// --> 1 <= Node.val <= 10^5

func nodesBetweenCriticalPoints(head *ListNode) []int {
	res := make([]int, 2)
	res[0] = -1
	res[1] = -1

	curr := head
	for {
		next := curr.Next
		if next == nil {
			return res
		}

		nextToNext := next.Next
		if nextToNext == nil {
			return res
		}

		if (curr.Val < next.Val && nextToNext.Val < next.Val) ||
			(curr.Val > next.Val && nextToNext.Val > next.Val) {
			curr = next
			break
		}
		curr = next
	}

	count := 1
	for {
		next := curr.Next
		if next == nil {
			return res
		}

		nextToNext := next.Next
		if nextToNext == nil {
			return res
		}

		if (curr.Val < next.Val && nextToNext.Val < next.Val) ||
			(curr.Val > next.Val && nextToNext.Val > next.Val) {
			curr = next
			res[0] = count
			res[1] = count
			count = 1
			break
		}
		curr = next
		count++
	}

	for {
		next := curr.Next
		if next == nil {
			break
		}

		nextToNext := next.Next
		if nextToNext == nil {
			break
		}

		if (curr.Val < next.Val && nextToNext.Val < next.Val) ||
			(curr.Val > next.Val && nextToNext.Val > next.Val) {
			res[0] = min(res[0], count)
			res[1] += count
			count = 1
		} else {
			count++
		}
		curr = next
	}

	return res
}
