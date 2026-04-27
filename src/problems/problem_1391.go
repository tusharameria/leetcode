// 1391. Check if There is a Valid Path in a Grid

package problems

import "fmt"

// 1 : L -> R
// 2 : U -> D
// 3 : L -> D
// 4 : R -> D
// 5 : L -> U
// 6 : R -> U

func Problem_1391() {
	grid := [][]int{{4, 3, 3}, {6, 5, 2}}
	fmt.Println(hasValidPath(grid))
}

func hasValidPath(grid [][]int) bool {
	return (len(grid) == 1 && len(grid[0]) == 1) || hasValidPathFunction(grid, 0, 0, 0)
}

func hasValidPathFunction(grid [][]int, x int, y int, direction int) bool {
	if x == 0 && y == 0 {
		if grid[0][0] == 1 || grid[0][0] == 6 {
			return hasValidPathFunction(grid, x+1, y, 0)
		} else if grid[0][0] == 2 || grid[0][0] == 3 {
			return hasValidPathFunction(grid, x, y+1, 2)
		} else if grid[0][0] == 4 {
			return hasValidPathFunction(grid, x+1, y, 0) || hasValidPathFunction(grid, x, y+1, 2)
		} else {
			return false
		}
	} else {
		width := len(grid[0])
		height := len(grid)
		memo := make([][]bool, height)
		for i := 0; i < height; i++ {
			memo[i] = make([]bool, width)
		}
		memo[0][0] = true
		for x >= 0 && y >= 0 && x <= width-1 && y <= height-1 {
			if memo[y][x] {
				return false
			}
			memo[y][x] = true
			xInit := x
			yInit := y
			val := grid[y][x]
			if direction == 0 {
				if val == 1 || val == 3 || val == 5 {
					if val == 1 {
						x++
					} else if val == 3 {
						y++
						direction = 2
					} else {
						y--
						direction = 3
					}
				} else {
					return false
				}
			} else if direction == 1 {
				if val == 1 || val == 4 || val == 6 {
					if val == 1 {
						x--
					} else if val == 4 {
						y++
						direction = 2
					} else {
						y--
						direction = 3
					}
				} else {
					return false
				}
			} else if direction == 2 {
				if val == 2 || val == 5 || val == 6 {
					if val == 2 {
						y++
					} else if val == 5 {
						x--
						direction = 1
					} else {
						x++
						direction = 0
					}
				} else {
					return false
				}
			} else {
				if val == 2 || val == 3 || val == 4 {
					if val == 2 {
						y--
					} else if val == 3 {
						x--
						direction = 1
					} else {
						x++
						direction = 0
					}
				} else {
					return false
				}
			}
			if xInit == width-1 && yInit == height-1 {
				return true
			}
		}
		return false
	}
}
