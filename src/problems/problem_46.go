// 46. Permutations

package problems

import "fmt"

func Problem_46() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}

func permute(nums []int) [][]int {
	res := [][]int{}

	var fill func([]int, []int)
	fill = func(initArray, leftArray []int) {
		if len(leftArray) == 0 {
			res = append(res, initArray)
		}
		for i := 0; i < len(leftArray); i++ {
			buffInit := make([]int, len(initArray))
			buffLeft := make([]int, len(leftArray))
			copy(buffInit, initArray)
			copy(buffLeft, leftArray)
			fill(append(buffInit, leftArray[i]), append(buffLeft[:i], buffLeft[i+1:]...))
		}
	}

	fill([]int{}, nums)
	return res
}
