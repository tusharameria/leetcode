// 3548. Equal Sum Grid Partition II

package problems

import (
	"fmt"
	"math"
)

func Problem_3548() {
	grid := [][]int{
		{54756, 54756},
	}
	fmt.Println(canPartitionGrid(grid))
}

type point struct {
	r, c int
}

func canRemove(r1, c1, r2, c2, i, j int) bool {
	rows := r2 - r1 + 1
	cols := c2 - c1 + 1

	if rows*cols <= 1 {
		return false
	}

	if rows == 1 {
		return j == c1 || j == c2
	}
	if cols == 1 {
		return i == r1 || i == r2
	}

	return true
}

func canPartitionGrid(grid [][]int) bool {
	n := len(grid)
	m := len(grid[0])

	prefRow := make([]int64, n)
	prefCol := make([]int64, m)
	mp := make(map[int][]point)

	for i := 0; i < n; i++ {
		var rowSum int64
		for j := 0; j < m; j++ {
			val := grid[i][j]
			rowSum += int64(val)
			mp[val] = append(mp[val], point{i, j})
		}
		if i > 0 {
			prefRow[i] = rowSum + prefRow[i-1]
		} else {
			prefRow[i] = rowSum
		}
	}

	for j := 0; j < m; j++ {
		var colSum int64
		for i := 0; i < n; i++ {
			colSum += int64(grid[i][j])
		}
		if j > 0 {
			prefCol[j] = colSum + prefCol[j-1]
		} else {
			prefCol[j] = colSum
		}
	}

	total := prefRow[n-1]

	for i := 0; i < n-1; i++ {
		top := prefRow[i]
		bottom := total - top

		if top == bottom {
			return true
		}

		diff := int(math.Abs(float64(top - bottom)))
		if pts, ok := mp[diff]; ok {
			for _, p := range pts {
				if top > bottom {

					if p.r <= i && canRemove(0, 0, i, m-1, p.r, p.c) {
						return true
					}
				} else {

					if p.r > i && canRemove(i+1, 0, n-1, m-1, p.r, p.c) {
						return true
					}
				}
			}
		}
	}

	for j := 0; j < m-1; j++ {
		left := prefCol[j]
		right := total - left

		if left == right {
			return true
		}

		diff := int(math.Abs(float64(left - right)))
		if pts, ok := mp[diff]; ok {
			for _, p := range pts {
				if left > right {

					if p.c <= j && canRemove(0, 0, n-1, j, p.r, p.c) {
						return true
					}
				} else {

					if p.c > j && canRemove(0, j+1, n-1, m-1, p.r, p.c) {
						return true
					}
				}
			}
		}
	}

	return false
}
