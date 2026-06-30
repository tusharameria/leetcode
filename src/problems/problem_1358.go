// 1358. Number of Substrings Containing All Three Characters

package problems

import "fmt"

func Problem_1358() {
	s := "ababbbc"
	fmt.Println(numberOfSubstrings(s))
}

func numberOfSubstrings(s string) int {
	fmt.Println(s)
	n := len(s)
	idxStore := [3]int{n, n, n}

	for i := n - 1; i >= 0; i-- {
		idxStore[int(s[i]-'a')] = i
	}

	left, mid, right := idxStore[0], idxStore[1], idxStore[2]

	if mid > right {
		mid, right = right, mid
	}

	if left > right {
		left, right = right, left
	}

	if left > mid {
		left, mid = mid, left
	}

	if right == n {
		return 0
	}

	res := 0

	for {
		fmt.Println("============")
		fmt.Printf("res : %d\n", res)
		fmt.Printf("left : %d\n", left)
		fmt.Printf("mid : %d\n", mid)
		fmt.Printf("right : %d\n", right)
		res += (n - right) * (mid - left)
		newIdx := n
		for i := mid + 1; i < n; i++ {
			if s[i] == s[left] {
				newIdx = i
				break
			}
		}
		if newIdx == n {
			return res
		}
		if newIdx < right {
			left = mid
			mid = newIdx
		} else {
			left, mid, right = mid, right, newIdx
		}
	}
}
