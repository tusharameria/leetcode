// 1861. Rotating the Box

package problems

import "fmt"

func Problem_1861() {
	box := [][]byte{
		{'#', '.', '*', '.'},
		{'#', '#', '*', '.'},
	}
	fmt.Println(rotateTheBox(box))
}

func rotateTheBox(boxGrid [][]byte) [][]byte {
	rows := len(boxGrid)
	cols := len(boxGrid[0])
	rotated := make([][]byte, cols)
	for i := range rotated {
		rotated[i] = make([]byte, rows)
		for j := range rotated[i] {
			rotated[i][j] = '.'
		}
	}
	for i := range rows {
		rotRow := cols - 1
		for j := cols - 1; j >= 0; j-- {
			if boxGrid[i][j] == '*' {
				rotated[j][rows-1-i] = '*'
				rotRow = j - 1
			} else if boxGrid[i][j] == '#' {
				rotated[rotRow][rows-1-i] = '#'
				rotRow--
			}
		}
	}
	return rotated
}
