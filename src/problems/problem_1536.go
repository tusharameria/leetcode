// 1536. Minimum Swaps to Arrange a Binary Grid

package problems

import (
	"fmt"
)

func Problem_1536() {
	// grid := utils.RandomInt2DArrayGenerator(0, 1, 4, 4)
	grid := [][]int{
		{1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0, 0},
		{1, 0, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{1, 1, 0, 1, 0, 0},
		{1, 0, 0, 0, 0, 0},
	}
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println(minSwaps(grid))
}

func minSwaps(grid [][]int) int {
	n := len(grid)
	res := 0

	zeroArr := make([]int, n)

	for i := 0; i < n; i++ {
		trailingZeroes := n
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] == 1 {
				trailingZeroes = n - (j + 1)
				break
			}
		}
		zeroArr[i] = trailingZeroes
	}

	for i := 0; i < n; i++ {
		required := (n - 1) - i
		if zeroArr[i] < required {
			targetIdx := n
			for k := i + 1; k < n; k++ {
				if zeroArr[k] >= required {
					targetIdx = k
					break
				}
			}
			if targetIdx == n {
				return -1
			}
			for k := targetIdx; k > i; k-- {
				zeroArr[k], zeroArr[k-1] = zeroArr[k-1], zeroArr[k]
			}
			res += targetIdx - i
		}
	}
	return res
}
