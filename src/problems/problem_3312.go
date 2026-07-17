// 3312. Sorted GCD Pair Queries

package problems

import (
	"fmt"
)

func Problem_3312() {
	nums := []int{3, 2, 4}
	queries := []int64{0, 1, 2}
	fmt.Println(gcdValues(nums, queries))
}

func gcdValues(nums []int, queries []int64) []int {

	mx := 0
	for _, x := range nums {
		if x > mx {
			mx = x
		}
	}

	freq := make([]int, mx+1)
	for _, x := range nums {
		freq[x]++
	}

	exact := make([]int64, mx+1)

	for g := mx; g >= 1; g-- {
		var cnt int64 = 0
		for m := g; m <= mx; m += g {
			cnt += int64(freq[m])
		}
		pairs := cnt * (cnt - 1) / 2
		for m := g * 2; m <= mx; m += g {
			pairs -= exact[m]
		}
		exact[g] = pairs
	}

	prefix := make([]int64, mx+1)
	for g := 1; g <= mx; g++ {
		prefix[g] = prefix[g-1] + exact[g]
	}

	ans := make([]int, len(queries))

	for i, q := range queries {
		l, r := 1, mx
		for l < r {
			mid := (l + r) / 2
			if prefix[mid] >= q+1 {
				r = mid
			} else {
				l = mid + 1
			}
		}
		ans[i] = l
	}

	return ans
}
