package problems

import "fmt"

func Problem_3721() {
	nums := []int{1, 2, 3, 2, 5}
	fmt.Println(longestBalanced2(nums))
}

func longestBalanced2(nums []int) int {
	n := len(nums)

	// Fenwick Tree
	bit := make([]int, n+2)

	add := func(i, v int) {
		for i <= n+1 {
			bit[i] += v
			i += i & -i
		}
	}

	sum := func(i int) int {
		s := 0
		for i > 0 {
			s += bit[i]
			i -= i & -i
		}
		return s
	}

	last := make(map[int]int)
	first := make(map[int]int)
	first[0] = 0

	ans := 0

	for i := 0; i < n; i++ {
		x := nums[i]
		val := 1
		if x%2 == 1 {
			val = -1
		}

		if j, ok := last[x]; ok {
			add(j+1, -val)
		}
		add(i+1, val)
		last[x] = i

		balance := sum(i + 1)

		if idx, ok := first[balance]; ok {
			if i+1-idx > ans {
				ans = i + 1 - idx
			}
		} else {
			first[balance] = i + 1
		}
	}

	return ans
}
