// 2657. Find the Prefix Common Array of Two Arrays

package problems

import "fmt"

func Problem_2657() {
	A := []int{1, 3, 2, 4}
	B := []int{3, 1, 2, 4}
	fmt.Println(findThePrefixCommonArray(A, B))
}

func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	if n == 1 {
		return []int{1}
	}
	aStore := make([]bool, n)
	bStore := make([]bool, n)
	res := make([]int, n)

	aStore[A[0]-1] = true
	bStore[B[0]-1] = true
	if A[0] == B[0] {
		res[0] = 1
	}
	for i := 1; i < n-1; i++ {
		aVal := A[i] - 1
		bVal := B[i] - 1
		if aVal == bVal {
			res[i] = res[i-1] + 1
		} else {
			inc := 0
			if aStore[bVal] {
				inc++
			}
			if bStore[aVal] {
				inc++
			}
			res[i] = res[i-1] + inc
		}
		aStore[aVal] = true
		bStore[bVal] = true
	}
	res[n-1] = n
	return res
}
