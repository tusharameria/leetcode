// 3346. Maximum Frequency of an Element After Performing Operations I

package problems

import "fmt"

func Problem_3346() {
	nums := []int{5, 11, 20, 20}
	k := 5
	numOperations := 1
	fmt.Println(maxFrequency(nums, k, numOperations))
}

func maxFrequency(nums []int, k int, numOperations int) int {
	min := nums[0]
	max := nums[0]
	startIndices := make(map[int]int)
	endIndices := make(map[int]int)
	numFreq := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
		startIndices[nums[i]-k]++
		endIndices[nums[i]+k]++
		numFreq[nums[i]]++
	}
	maxFreq := 0
	freq := 0
	// fmt.Println("startIndices : ", startIndices)
	// fmt.Println("endIndices : ", endIndices)
	for i := range max - min + 1 + (2 * k) {
		number := i + (min - k)
		// fmt.Println("======================")
		// fmt.Println("number : ", number)
		freq += startIndices[number]
		// fmt.Println("freq : ", freq)
		testingFreq := freq - numFreq[number]
		if testingFreq <= numOperations {
			if freq > maxFreq {
				maxFreq = freq
			}
		} else {
			if numOperations+numFreq[number] > maxFreq {
				maxFreq = numOperations + numFreq[number]
			}
		}
		freq -= endIndices[number]
	}
	return maxFreq
}
