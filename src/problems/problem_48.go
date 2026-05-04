// 48. Rotate Image

package problems

import "fmt"

func Problem_48() {
	matrix := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(matrix)
	fmt.Println("res : ", matrix)
}

func rotate(matrix [][]int) {
	matLen := len(matrix)
	for i := 0; i < matLen/2; i++ {
		for j := i; j < matLen-1-i; j++ {
			// Process to rotate four elements in clockwise direction
			buff := j - i
			temp := matrix[i][j]
			matrix[i][j] = matrix[matLen-1-i-buff][i]
			matrix[matLen-1-i-buff][i] = matrix[matLen-1-i][matLen-1-i-buff]
			matrix[matLen-1-i][matLen-1-i-buff] = matrix[i+buff][matLen-1-i]
			matrix[i+buff][matLen-1-i] = temp
		}
	}
}
