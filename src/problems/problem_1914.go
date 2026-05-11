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
	res := make([][]int, height)
	for i := 0; i < height; i++ {
		res[i] = make([]int, width)
		copy(res[i], grid[i])
	}
	layers := min(height, width) / 2
	for layer := 0; layer < layers; layer++ {
		// fmt.Printf("Layer : %d\n", layer)
		layerHeight := height - 2*layer
		layerWidth := width - 2*layer
		layerElements := 2 * (layerHeight + layerWidth - 2)
		// fmt.Printf("layerHeight : %d\n", layerHeight)
		// fmt.Printf("layerWidth : %d\n", layerWidth)
		// fmt.Printf("layerElements : %d\n", layerElements)
		shift := k % layerElements
		// fmt.Printf("Shift : %d\n", shift)
		if shift != 0 {
			// shift range => [1, layerElements-1]
			for count := 0; count < layerElements; count++ {
				nextCount := (count + layerElements - shift) % layerElements
				currentRow := 0
				currentCol := 0
				if count < layerWidth {
					currentCol = count
				} else if count < layerWidth+layerHeight-1 {
					currentRow = count - layerWidth + 1
					currentCol = layerWidth - 1
				} else if count < 2*layerWidth+layerHeight-2 {
					currentRow = layerHeight - 1
					currentCol = 2*layerWidth + layerHeight - 3 - count
				} else {
					currentRow = layerHeight - 1 - (count - (2*layerWidth + layerHeight - 3))
				}
				currentRow += layer
				currentCol += layer
				// For next element
				nextRow := 0
				nextCol := 0
				if nextCount < layerWidth {
					nextCol = nextCount
				} else if nextCount < layerWidth+layerHeight-1 {
					nextRow = nextCount - layerWidth + 1
					nextCol = layerWidth - 1
				} else if nextCount < 2*layerWidth+layerHeight-2 {
					nextRow = layerHeight - 1
					nextCol = 2*layerWidth + layerHeight - 3 - nextCount
				} else {
					nextRow = layerHeight - 1 - (nextCount - (2*layerWidth + layerHeight - 3))
				}
				nextRow += layer
				nextCol += layer
				res[nextRow][nextCol] = grid[currentRow][currentCol]
			}
		}
	}
	return res
}
