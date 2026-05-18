// 1306. Jump Game III

package problems

import "fmt"

func Problem_1306() {
	arr := []int{4, 2, 3, 0, 3, 1, 2}
	start := 5
	fmt.Println(canReach(arr, start))
}

func canReach(arr []int, start int) bool {
	n := len(arr)
	store := make([][2]bool, n)

	zero := false
	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			zero = true
			break
		}
	}

	if !zero {
		return false
	}

	var iter func(index int) bool

	iter = func(index int) bool {
		val := arr[index]
		if val == 0 {
			return true
		}
		// both side already visited
		if store[index][0] && store[index][1] {
			return false
		}
		leftIndex := index - val
		rightIndex := index + val
		// both side out of bounds
		if leftIndex < 0 && rightIndex >= n {
			store[index][0] = true
			store[index][1] = true
			return false
		}
		leftPossible := !store[index][0] && leftIndex >= 0
		rightPossible := !store[index][1] && rightIndex < n
		store[index][0] = true
		store[index][1] = true
		if leftPossible && !rightPossible {
			return iter(leftIndex)
		} else if !leftPossible && rightPossible {
			return iter(rightIndex)
		} else {
			return iter(leftIndex) || iter(rightIndex)
		}
	}

	return iter(start)
}
