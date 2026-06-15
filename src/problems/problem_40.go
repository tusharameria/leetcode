// 40. Combination Sum II

// Constraints:
// 1 <= candidates.length <= 30
// 2 <= candidates[i] <= 40
// All elements of candidates are distinct.
// 1 <= target <= 40

package problems

import (
	"fmt"
)

func Problem_40() {
	candidates := []int{10, 1, 2, 6, 1}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}

	var dfs func(int, int, []int)

	dfs = func(start, newTarget int, curr []int) {
		if newTarget <= 0 {
			if newTarget == 0 {
				if newTarget == 0 {
					buff := make([]int, len(curr))
					copy(buff, curr)
					res = append(res, buff)
				}
			}
			return
		}

		for i := start; i < len(candidates); i++ {
			dfs(i+1, newTarget-candidates[i], append(curr, candidates[i]))
		}
	}

	dfs(0, target, []int{})
	return res
}
