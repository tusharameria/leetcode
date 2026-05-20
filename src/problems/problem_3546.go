// 3546. Equal Sum Grid Partition I

package problems

import "fmt"

func Problem_3546() {
	grid := [][]int{
		{54756, 54756},
	}
	fmt.Println(canPartitionGrid(grid))
}

func canPartitionGrid(grid [][]int) bool {
	height := len(grid)
	width := len(grid[0])
	if height == 1 && width == 1 {
		return false
	}
	rowSums := make([]int, height)
	for j := 0; j < width; j++ {
		rowSums[0] += grid[0][j]
	}
	for i := 1; i < height; i++ {
		for j := 0; j < width; j++ {
			rowSums[i] += grid[i][j]
		}
		rowSums[i] += rowSums[i-1]
	}
	sum := rowSums[height-1]
	if sum%2 != 0 {
		return false
	}
	colSums := make([]int, width)
	for i := 0; i < height; i++ {
		colSums[0] += grid[i][0]
	}
	if colSums[0]*2 == sum {
		return true
	}
	for j := 1; j < width-1; j++ {
		for i := 0; i < height; i++ {
			colSums[j] += grid[i][j]
		}
		colSums[j] += colSums[j-1]
		if colSums[j]*2 == sum {
			return true
		}
	}
	for i := 0; i < height-1; i++ {
		if rowSums[i]*2 == sum {
			return true
		}
	}
	return false
}
