package problems

import "fmt"

// 3666. Minimum Operations to Equalize Binary String

func Problem_3666() {
	s := "0101"
	k := 3
	fmt.Println("res : ", minOperations(s, k))
}

func minOperations(s string, k int) int {
	zero := 0
	one := 0
	totalCount := len(s)
	for i := 0; i <= totalCount-1; i++ {
		if s[i] == '1' {
			one++
		} else {
			zero++
		}
	}
	if k == 1 {
		return zero
	}
	if zero == 0 {
		return 0
	}
	if k%2 == 0 && zero%2 == 1 {
		return -1
	}
	if totalCount == k && zero%k != 0 {
		return -1
	}
	ans := 0
	base := totalCount - k
	if (k % 2) == (zero % 2) {
		m := max(
			(zero+k-1)/k,
			(one+base-1)/base,
		)
		if m%2 == 0 {
			m++
		}
		if m < ans {
			ans = m
		}
	}

	if zero%2 == 0 {
		m := max(
			(zero+k-1)/k,
			(zero+base-1)/base,
		)
		if m%2 == 1 {
			m++
		}
		if m < ans {
			ans = m
		}
	}

	if ans == int(^uint(0)>>1) {
		return -1
	}
	return ans
}
