// 39. Combination Sum

// Constraints:
// 1 <= candidates.length <= 30
// 2 <= candidates[i] <= 40
// All elements of candidates are distinct.
// 1 <= target <= 40

package problems

import (
	"fmt"
)

func Problem_39() {
	candidates := []int{2, 3, 5}
	target := 8
	fmt.Println(combinationSum(candidates, target))
}

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var backtrack func([]int, int, []int)

	backtrack = func(candidates []int, target int, selected []int) {
		if target == 0 {
			result = append(result, append([]int{}, selected...))
			return
		}
		for i, val := range candidates {
			if val <= target {
				newTarget := target - val
				selected := append([]int{}, selected...)
				selected = append(selected, val)
				newCandidates := append([]int{}, candidates[i:]...)
				backtrack(newCandidates, newTarget, selected)
			}
		}
	}
	backtrack(candidates, target, []int{})
	return result
}
