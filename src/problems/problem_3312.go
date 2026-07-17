// 3312. Sorted GCD Pair Queries

package problems

import (
	"fmt"
	"sort"
)

func Problem_3312() {
	nums := []int{3, 2, 4}
	queries := []int64{0, 1, 2}
	fmt.Println(gcdValues(nums, queries))
}

const LUT_NEW_SIZE = 64

var LUT_NEW [64][64]uint8

func init() {
	for i := uint8(1); i < LUT_NEW_SIZE; i++ {
		LUT_NEW[i][i], LUT_NEW[0][i], LUT_NEW[i][0] = i, i, i
		for j := i + 1; j < LUT_NEW_SIZE; j++ {
			r := LUT_NEW[i][j-i]
			LUT_NEW[i][j], LUT_NEW[j][i] = r, r
		}
	}
	// for _, row := range LUT_NEW {
	// 	for _, val := range row {
	// 		fmt.Printf("%0d ", val)
	// 	}
	// 	fmt.Println()
	// }
}

func gcdValues(nums []int, queries []int64) []int {
	queryLen := len(queries)
	n := len(nums)

	gcdPairs := make([]int, int64((n*(n-1))/2))

	for i := 0; i < n; i++ {
		a := nums[i]
		for j := i + 1; j < n; j++ {
			b := nums[j]
			fmt.Printf("a : %d, b : %d\n", a, b)
			fmt.Printf("gcd(a, b) : %d\n", gcd(a, b))
			gcdPairs[(i*n)+j-((i+1)*(i+2)/2)] = gcd(a, b)
		}
	}

	sort.Ints(gcdPairs)
	res := make([]int, queryLen)
	for i := 0; i < queryLen; i++ {
		res[i] = gcdPairs[queries[i]]
	}
	return res
}

// a is smaller
func gcd(a, b int) int {
	if b < a {
		a, b = b, a
	}
	for a >= LUT_NEW_SIZE {
		a, b = b%a, a
	}
	if a == 0 {
		return b
	}
	return int(LUT_NEW[b%a][a])
}
