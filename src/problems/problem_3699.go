// 3699. Number of ZigZag Arrays I

package problems

import (
	"fmt"
)

func Problem_3699() {
	n := 8183787
	l := 21
	r := 27
	fmt.Println(zigZagArrays(n, l, r))
}

func zigZagArrays(n int, l int, r int) int {
	nums := r - l + 1
	if nums == 2 {
		return 2
	}

	store := make([]int, nums)
	buff := make([]int, nums)

	for i := 0; i < nums; i++ {
		store[i] = i
	}

	for i := 2; i < n; i++ {
		copy(buff, store)
		for j := nums - 2; j > 0; j-- {
			buff[j] += buff[j+1] % 1000000007
		}
		for j := 1; j < nums; j++ {
			store[j] = buff[nums-j] % 1000000007
		}
	}

	sum := 0
	for i := range nums {
		sum += store[i]
	}

	return (sum * 2) % 1000000007
}
