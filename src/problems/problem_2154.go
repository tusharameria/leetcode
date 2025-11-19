package problems

import (
	"fmt"
	"sort"
)

func Problem_2154() {
	nums := []int{2, 4, 6, 7, 8}
	original := 2
	fmt.Println(findFinalValue(nums, original))
}

func findFinalValue(nums []int, original int) int {
	sort.Ints(nums)
	for i := 0; i <= len(nums)-1; i++ {
		if original == nums[i] {
			original *= 2
		}
	}
	return original
}

// func findFinalValue(nums []int, original int) int {
// 	for j := 0; j <= len(nums)-1; j++ {

// 		foundOriginal := false
// 		for i := 0; i <= len(nums)-1; i++ {
// 			if original == nums[i] {
// 				original = 2 * original
// 				foundOriginal = true
// 				break
// 			}
// 		}
// 		if !foundOriginal {
// 			return original
// 		}
// 	}
// 	return original
// }
