// 1345. Jump Game IV

package problems

import "fmt"

func Problem_1345() {
	nums := []int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}
	fmt.Println(minJumps(nums))
}

func minJumps(arr []int) int {
	n := len(arr)
	if n <= 2 {
		return n - 1
	}
	memory := make(map[int][]int)
	for index, val := range arr {
		memory[val] = append(memory[val], index)
	}
	checked := make([]bool, 50000)
	queue := make([]int, 1)
	queue[0] = 0
	steps := 0
	checked[0] = true
	for len(queue) > 0 {
		size := len(queue)
		for s := 0; s < size; s++ {
			currentIndex := queue[0]
			queue = queue[1:]
			if currentIndex == n-1 {
				return steps
			}
			if currentIndex-1 > 0 && !checked[currentIndex-1] {
				checked[currentIndex-1] = true
				queue = append(queue, currentIndex-1)
			}
			if currentIndex+1 < n-1 && !checked[currentIndex+1] {
				checked[currentIndex+1] = true
				queue = append(queue, currentIndex+1)
			}
			val := arr[currentIndex]
			if len(memory[val]) > 1 {
				for _, index := range memory[val] {
					if !checked[index] {
						checked[index] = true
						queue = append(queue, index)
					}
				}
			}
			delete(memory, val)
		}
		steps++
	}
	return steps
}
