// 2033. Minimum Operations to Make a Uni-Value Grid

package problems

import (
	"fmt"
	"sort"
)

func Problem_2033() {
	grid := [][]int{{2, 4}, {6, 8}}
	fmt.Println(minOperations(grid, 2))
}

func minOperations(grid [][]int, x int) int {
	height := len(grid)
	width := len(grid[0])
	nums := make([]int, height*width)
	rem := grid[0][0] % x
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j]%x != rem {
				return -1
			}
			nums[(i*width)+j] = grid[i][j]
		}
	}
	sort.Ints(nums)
	mid := (len(nums) + 1) / 2
	midVal := nums[mid]
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans += abs(nums[i]-midVal) / x
	}
	return ans
}
