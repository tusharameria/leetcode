// 46. Permutations

package problems

import "fmt"

func Problem_46() {
	nums := []int{1, 1, 2}
	fmt.Println(permute(nums))
}

func permute(nums []int) [][]int {
	// n := len(nums)
	// resLen := 1
	// for i := 2; i <= n; i++ {
	// 	resLen *= i
	// }
	res := [][]int{}

	idx := 0
	var fill func([]int, []int)
	fill = func(initArray, leftArray []int) {
		leftLen := len(leftArray)
		if leftLen == 0 {
			res = append(res, initArray)
			idx++
		}
		for i := 0; i < leftLen; i++ {
			buffInit := make([]int, len(initArray))
			buffLeft := make([]int, leftLen)
			copy(buffInit, initArray)
			copy(buffLeft, leftArray)
			fill(append(buffInit, leftArray[i]), append(buffLeft[:i], buffLeft[i+1:]...))
		}
	}

	fill([]int{}, nums)
	return res
}
