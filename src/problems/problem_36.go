package problems

import "fmt"

// 36. Valid Sudoku

func Problem_36() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	fmt.Println("res : ", isValidSudoku(board))
}

func isValidSudoku(board [][]byte) bool {

	for i := 0; i < 9; i++ {
		rowSet := make(map[byte]int)
		colSet := make(map[byte]int)
		boxSet := make(map[byte]int)
		for j := 0; j < 9; j++ {
			row := i
			col := j
			if board[row][col] != '.' {
				rowSet[board[row][col]]++
				if rowSet[board[row][col]] > 1 {
					return false
				}
			}
			row = j
			col = i
			if board[row][col] != '.' {
				colSet[board[row][col]]++
				if colSet[board[row][col]] > 1 {
					return false
				}
			}
			row = (i/3)*3 + (j / 3)
			col = (i%3)*3 + (j % 3)
			if board[row][col] != '.' {
				boxSet[board[row][col]]++
				if boxSet[board[row][col]] > 1 {
					return false
				}
			}
		}
	}

	return true
}
