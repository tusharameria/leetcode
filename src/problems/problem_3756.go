// 3756. Concatenate Non-Zero Digits and Multiply by Sum II

package problems

import (
	"fmt"
)

func Problem_3756() {
	s := "10203004"
	queries := [][]int{
		{0, 7},
		{1, 3},
		{4, 6},
	}
	fmt.Println(sumAndMultiply(s, queries))
}

const MAXN = 100_000

var pow10, xs [MAXN + 1]int
var sums, digits [MAXN + 1]int

func init() {
	pow10[0] = 1
	for i := 1; i <= MAXN; i++ {
		pow10[i] = (pow10[i-1] * 10) % MOD
	}
}

func sumAndMultiply(s string, queries [][]int) []int {
	n := len(s)
	queryLen := len(queries)

	for i := range n {
		d := int(s[i] - '0')
		sums[i+1] = sums[i] + d
		if d > 0 {
			xs[i+1] = ((xs[i] * 10) + d) % MOD
			digits[i+1] = digits[i] + 1
		} else {
			xs[i+1] = xs[i]
			digits[i+1] = digits[i]
		}
	}

	res := make([]int, queryLen)
	for i := range queryLen {
		l := queries[i][0]
		r := queries[i][1]
		length := digits[l] - digits[r]
		x := (xs[r] - xs[l]*pow10[length] + MOD*MOD) % MOD
		sum := int(sums[r] - sums[l])
		res[i] = (x * sum) % MOD
		// res[i] =
	}
	return res
}
