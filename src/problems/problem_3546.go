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
	sum := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			sum += grid[i][j]
		}
	}

	currSum := 0

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			currSum += grid[i][j]
		}
		if currSum*2 >= sum {
			if currSum*2 == sum {
				return true
			}
			currSum = 0
			break
		}
	}

	for j := 0; j < width; j++ {
		for i := 0; i < height; i++ {
			currSum += grid[i][j]
		}
		if currSum*2 >= sum {
			if currSum*2 == sum {
				return true
			}
			return false
		}
	}

	return false
}
