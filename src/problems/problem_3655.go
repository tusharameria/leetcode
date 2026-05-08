// 3655. XOR After Range Multiplication Queries II

package problems

import (
	"fmt"
	"math"
)

func Problem_3655() {
	nums := []int{492, 903, 451, 625, 787}
	queries := [][]int{
		{3, 3, 1, 4},
		{2, 2, 3, 18},
		{2, 3, 2, 10},
		{2, 3, 5, 12},
		{1, 4, 3, 1},
		{3, 3, 3, 12},
		{3, 3, 4, 3},
		{0, 2, 1, 1},
		{2, 3, 5, 1},
		{4, 4, 2, 17},
		{0, 0, 2, 18},
	}
	fmt.Println(xorAfterQueries(nums, queries))
}

func xorAfterQueries(nums []int, queries [][]int) int {
	res := 0
	divider := int64(1000000007)
	queryMap := make(map[int][][3]int)
	n := len(nums)
	blockSize := int(math.Ceil(math.Sqrt(float64(n))))
	for i := range queries {
		l, r, step, multiplier := queries[i][0], queries[i][1], queries[i][2], queries[i][3]
		if step < blockSize {
			queryMap[step] = append(queryMap[step], [3]int{l, r, multiplier})
		}
	}
	for i := range queries {
		l, r, step, multiplier := queries[i][0], queries[i][1], queries[i][2], queries[i][3]
		if step >= blockSize {
			for j := l; j <= r; j += step {
				nums[j] = int((int64(nums[j]) * int64(multiplier)) % divider)
			}
		}
	}
	for step, queries := range queryMap {
		diff := make([]int64, n)
		for i := range diff {
			diff[i] = 1
		}
		for _, query := range queries {
			l, r, multiplier := query[0], query[1], query[2]
			diff[l] = (diff[l] * int64(multiplier)) % int64(divider)
			last := l + ((r-l)/step)*step
			if last+step < n {
				diff[last+step] =
					(diff[last+step] * modInverse(int64(multiplier))) % divider
			}
		}
		for rem := 0; rem < step; rem++ {
			for i := rem + step; i < n; i += step {
				diff[i] = (diff[i] * diff[i-step]) % divider
			}
		}
		for i := 0; i < n; i++ {
			nums[i] = int((int64(nums[i]) * diff[i]) % int64(divider))
		}
	}
	for _, v := range nums {
		res ^= v
	}
	return res
}

func modPow(a, b int64) int64 {
	res := int64(1)
	MOD := int64(1000000007)
	for b > 0 {
		if b&1 == 1 {
			res = (res * a) % MOD
		}
		a = (a * a) % MOD
		b >>= 1
	}

	return res
}

func modInverse(x int64) int64 {
	MOD := int64(1000000007)
	return modPow(x, MOD-2)
}
