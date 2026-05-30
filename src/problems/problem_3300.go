// 3300. Minimum Element After Replacement With Digit Sum

package problems

import "fmt"

func Problem_3300() {
	nums := []int{999, 19, 199}
	fmt.Println(minElement(nums))
}

func minElement(nums []int) int {
	res := 0
	buff := nums[0]
	for buff > 0 {
		res += buff % 10
		buff /= 10
	}
	if res == 1 {
		return 1
	}
	for i := 1; i < len(nums); i++ {
		currSum := 0
		num := nums[i]
		for num > 0 {
			currSum += num % 10
			num /= 10
		}
		res = min(res, currSum)
	}
	return res
}
