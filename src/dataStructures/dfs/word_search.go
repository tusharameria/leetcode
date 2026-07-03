package dfs

import "fmt"

func WordSearch() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCDEDF"
	fmt.Println(exist(board, word))
}

func exist(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])
	wordLen := len(word)
	if wordLen > rows*cols {
		return false
	}
	count := 0
	startCoord := make([][2]int, rows*cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if word[0] == board[i][j] {
				startCoord[count] = [2]int{i, j}
				count++
			}
		}
	}
	startCoord = startCoord[:count]

	if count == 0 {
		return false
	}
	if wordLen == 1 {
		return true
	}

	visited := make([][]bool, rows)
	for i := range rows {
		visited[i] = make([]bool, cols)
	}

	var iterateBoard func(x, y, index int) bool
	iterateBoard = func(x, y, index int) bool {
		if index == wordLen {
			return true
		}

		target := word[index]

		left := []int{x - 1, y}
		up := []int{x, y - 1}
		right := []int{x + 1, y}
		down := []int{x, y + 1}

		if left[0] >= 0 && !visited[left[0]][left[1]] && board[left[0]][left[1]] == target {
			visited[left[0]][left[1]] = true
			if iterateBoard(left[0], left[1], index+1) {
				return true
			}
			visited[left[0]][left[1]] = false
		}

		if up[1] >= 0 && !visited[up[0]][up[1]] && board[up[0]][up[1]] == target {
			visited[up[0]][up[1]] = true
			if iterateBoard(up[0], up[1], index+1) {
				return true
			}
			visited[up[0]][up[1]] = false
		}

		if right[0] < rows && !visited[right[0]][right[1]] && board[right[0]][right[1]] == target {
			visited[right[0]][right[1]] = true
			if iterateBoard(right[0], right[1], index+1) {
				return true
			}
			visited[right[0]][right[1]] = false
		}

		if down[1] < cols && !visited[down[0]][down[1]] && board[down[0]][down[1]] == target {
			visited[down[0]][down[1]] = true
			if iterateBoard(down[0], down[1], index+1) {
				return true
			}
			visited[down[0]][down[1]] = false
		}

		return false
	}

	for i := 0; i < count; i++ {
		x, y := startCoord[i][0], startCoord[i][1]
		visited[x][y] = true
		if iterateBoard(x, y, 1) {
			return true
		}
		visited[x][y] = false
	}
	return false
}
