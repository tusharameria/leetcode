package problems

import "fmt"

func Problem_2872() {
	n := 5
	edges := [][]int{
		{0, 2},
		{1, 2},
		{1, 3},
		{2, 4},
	}
	values := []int{1, 8, 1, 4, 4}
	k := 6
	fmt.Println(maxKDivisibleComponents(n, edges, values, k))
}

func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	freqMap := make(map[int][]int, n)
	ans := 0
	for i := 0; i <= n-2; i++ {
		freqMap[edges[i][0]] = append(freqMap[edges[i][0]], edges[i][1])
		freqMap[edges[i][1]] = append(freqMap[edges[i][1]], edges[i][0])
	}
	for nodeVal, arr := range freqMap {
		if len(arr) == 1 && values[nodeVal]%k == 0 {
			ans++

		}
	}
	return ans
}
