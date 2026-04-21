// 3488. Closest Equal Element Queries

package problems

import "fmt"

func Problem_3488() {
	nums := []int{1, 3, 1, 4, 1, 3, 2}
	queries := []int{0, 3, 5}
	fmt.Println(solveQueries(nums, queries))
}

func solveQueries(nums []int, queries []int) []int {
	res := make([]int, len(queries))
	freqMap := make(map[int][]int)
	posIndex := make([]int, len(nums))
	for i := 0; i <= len(nums)-1; i++ {
		freqMap[nums[i]] = append(freqMap[nums[i]], i)
		posIndex[i] = len(freqMap[nums[i]]) - 1
	}

	for i := 0; i <= len(queries)-1; i++ {
		if len(freqMap[nums[queries[i]]]) > 1 {
			currentIndex := posIndex[queries[i]]
			if currentIndex == -1 {
				res[i] = -1
			} else {
				freqLen := len(freqMap[nums[queries[i]]])
				prevIndex := (currentIndex - 1 + freqLen) % freqLen
				nextIndex := (currentIndex + 1) % freqLen
				currentVal := freqMap[nums[queries[i]]][currentIndex]
				prevVal := freqMap[nums[queries[i]]][prevIndex]
				nextVal := freqMap[nums[queries[i]]][nextIndex]
				res[i] = min(min(abs(currentVal-prevVal), len(nums)-abs(currentVal-prevVal)), min(abs(currentVal-nextVal), len(nums)-abs(currentVal-nextVal)))
			}
		} else {
			res[i] = -1
		}
	}
	return res
}
