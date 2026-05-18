// 1306. Jump Game III

package problems

import "fmt"

func Problem_1306() {
	arr := []int{4, 2, 3, 0, 3, 1, 2}
	start := 5
	fmt.Println(canReach(arr, start))
}

func canReach(arr []int, start int) bool {
	const SENTINEL = 99999999
	n := len(arr)
	store := make([]int, 1)
	store[0] = start
	for len(store) > 0 {
		index := store[0]
		val := arr[index]
		if val == 0 {
			return true
		}
		store = store[1:]
		p := []int{index - val, index + val}
		arr[index] = SENTINEL
		for _, newIndex := range p {
			if newIndex >= 0 && newIndex < n {
				if arr[newIndex] == 0 {
					return true
				}
				if arr[newIndex] != SENTINEL {
					store = append(store, newIndex)
				}
			}
		}
	}
	return false
}
