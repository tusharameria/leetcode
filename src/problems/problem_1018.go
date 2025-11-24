package problems

import "fmt"

func Problem_1018() {
	nums := []int{1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0}
	fmt.Println(prefixesDivBy5(nums))
}

func prefixesDivBy5(nums []int) []bool {
	n := len(nums)
	res := make([]bool, n)
	sum := 0
	for i := 0; i <= n-1; i++ {
		sum = (sum*2 + nums[i]) % 5
		res[i] = sum == 0
	}
	return res
}
