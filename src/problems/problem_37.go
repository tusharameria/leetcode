package problems

// 37. Sudoku Solver

func Problem_37() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
}

func solveSudoku(board [][]byte) {
	for i := 0; i < 9; i++ {
		rowValues := make([]int, 9)
		colValues := make([]int, 9)
		boxValues := make([]int, 9)
		for j := 0; j < 9; j++ {
			row := i
			col := j
			if board[row][col] != '.' {
				idx := int(board[row][col] - '1')
				rowValues[idx] = 1
			}
			row = j
			col = i
			if board[row][col] != '.' {
				idx := int(board[row][col] - '1')
				colValues[idx] = 1
			}
			row = (i/3)*3 + (j / 3)
			col = (i%3)*3 + (j % 3)
			if board[row][col] != '.' {
				idx := int(board[row][col] - '1')
				boxValues[idx] = 1
			}
		}

		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for k := 0; k < 9; k++ {
					if rowValues[k] == 0 && colValues[k] == 0 && boxValues[k] == 0 {
					}
				}
			}
		}
	}
}
