// 3488. Closest Equal Element Queries

package problems

import "fmt"

func Problem_3488() {
	nums := []int{1, 3, 1, 4, 1, 3, 2}
	queries := []int{0, 3, 5}
	fmt.Println(solveQueries(nums, queries))
}

func solveQueries(nums []int, queries []int) []int {
	freqMap := make(map[int][]int)
	for i := 0; i <= len(nums)-1; i++ {
		freqMap[nums[i]] = append(freqMap[nums[i]], i)
	}

	res := make([]int, len(queries))
	for i := 0; i <= len(queries)-1; i++ {
		if len(freqMap[nums[queries[i]]]) > 1 {
			currentIndex := findIndex(freqMap[nums[queries[i]]], queries[i])
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

func findIndex(arr []int, target int) int {
	start := 0
	end := len(arr) - 1
	for start <= end {
		mid := (start + end) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
