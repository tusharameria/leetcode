package problems

import "fmt"

// 799. Champagne Tower

func Problem_799() {
	poured := 25
	query_row := 6
	query_glass := 1
	fmt.Println("res : ", champagneTower(poured, query_row, query_glass))
}

func champagneTower(poured int, query_row int, query_glass int) float64 {
	if poured == 0 {
		return 0
	}
	var dp [101][101]float64
	dp[0][0] = float64(poured)

	for r := 0; r <= query_row; r++ {
		for c := 0; c <= r; c++ {
			overflow := (dp[r][c] - 1) / 2
			if overflow > 0 {
				dp[r+1][c] += overflow
				dp[r+1][c+1] += overflow
			}
		}
	}
	return min(dp[query_row][query_glass], 1)
}
