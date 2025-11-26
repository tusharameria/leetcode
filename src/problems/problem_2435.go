package problems

import "fmt"

func Problem_2435() {
	grid := [][]int{
		{1, 5, 3, 7, 3, 2, 3, 5},
	}
	k := 29
	fmt.Println(numberOfPaths(grid, k))
}

func numberOfPaths(grid [][]int, k int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][][]int, m)
	mod := 1_000_000_007

	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, k)
		}
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dp[i][j][grid[i][j]%k] = 1
			} else if i == m-1 && j != n-1 {
				for a := 0; a <= len(dp[i][j+1])-1; a++ {
					if dp[i][j+1][a] != 0 {
						dp[i][j][(grid[i][j]%k+a)%k] = 1
					}
				}
			} else if i != m-1 && j == n-1 {
				for a := 0; a <= len(dp[i+1][j])-1; a++ {
					if dp[i+1][j][a] != 0 {
						dp[i][j][(grid[i][j]%k+a)%k] = 1
					}
				}
			} else {
				for a := 0; a <= len(dp[i][j+1])-1; a++ {
					if dp[i][j+1][a] != 0 {
						dp[i][j][(grid[i][j]%k+a)%k] = (dp[i][j][(grid[i][j]%k+a)%k] + dp[i][j+1][a]) % mod
					}
				}
				for a := 0; a <= len(dp[i+1][j])-1; a++ {
					if dp[i+1][j][a] != 0 {
						dp[i][j][(grid[i][j]%k+a)%k] = (dp[i][j][(grid[i][j]%k+a)%k] + dp[i+1][j][a]) % mod
					}
				}
			}
		}
	}

	return dp[0][0][0]
}
