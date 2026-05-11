// 1914. Cyclically Rotating a Grid

package problems

import "fmt"

func Problem_1914() {
	grid := [][]int{
		{40, 10},
		{70, 50},
	}
	k := 3
	for _, row := range rotateGrid(grid, k) {
		fmt.Println(row)
	}
}

func rotateGrid(grid [][]int, k int) [][]int {
	height := len(grid)
	width := len(grid[0])
	layers := min(height, width) / 2
	for layer := 0; layer < layers; layer++ {
		rowIndex := []int{}
		colIndex := []int{}
		values := []int{}
		for i := layer; i <= width-layer-2; i++ {
			rowIndex = append(rowIndex, layer)
			colIndex = append(colIndex, i)
			values = append(values, grid[layer][i])
		}
		for i := layer; i <= height-layer-2; i++ {
			rowIndex = append(rowIndex, i)
			colIndex = append(colIndex, width-layer-1)
			values = append(values, grid[i][width-layer-1])
		}
		for i := width - layer - 1; i >= layer+1; i-- {
			rowIndex = append(rowIndex, height-layer-1)
			colIndex = append(colIndex, i)
			values = append(values, grid[height-layer-1][i])
		}
		for i := height - layer - 1; i >= layer+1; i-- {
			rowIndex = append(rowIndex, i)
			colIndex = append(colIndex, layer)
			values = append(values, grid[i][layer])
		}
		n := len(values)
		for i := 0; i < n; i++ {
			newIndex := (i + k%n) % n
			grid[rowIndex[i]][colIndex[i]] = values[newIndex]
		}
	}
	return grid
}
