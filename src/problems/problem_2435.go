package problems

import "fmt"

func Problem_2435() {
	grid := [][]int{
		{5, 2, 4},
		{3, 0, 5},
		{0, 7, 2},
	}
	k := 3
	fmt.Println(numberOfPaths(grid, k))
}

func numberOfPaths(grid [][]int, k int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, n)
	prev := make([][]int, n)
	mod := 1_000_000_007

	for i := range dp {
		dp[i] = make([]int, k)
		prev[i] = make([]int, k)
	}

	for i := m - 1; i >= 0; i-- {
		fmt.Printf("============== i : %d ==============\n", i)
		for j := n - 1; j >= 0; j-- {
			fmt.Printf("======= j : %d =======\n", j)
			fmt.Println("dp : ", dp)
			fmt.Println("prev : ", prev)
			val := grid[i][j]
			if i == m-1 && j == n-1 {
				prev[j][val%k] = 1
				dp[j][val%k] = 1
			} else if i == m-1 && j != n-1 {
				for a := 0; a <= k-1; a++ {
					if prev[j+1][a] != 0 {
						prev[j][(val+a)%k] = 1
						dp[j][(val+a)%k] = 1
					}
				}
			} else if i != m-1 && j == n-1 {
				for a := 0; a <= k-1; a++ {
					if prev[j][a] != 0 {
						dp[j][(val+a)%k] = 1
					}
				}
			} else {
				for a := 0; a <= k-1; a++ {
					if dp[j+1][a] != 0 {
						dp[j][(val+a)%k] = (dp[j+1][a]) % mod
					}
				}
				for a := 0; a <= k-1; a++ {
					if prev[j][a] != 0 {
						dp[j][(val+a)%k] = (dp[j][(val+a)%k] + prev[j][a]) % mod
					}
				}
			}
		}
		for a := range dp {
			prev[a] = dp[a]
			if i != 0 {
				dp[a] = make([]int, k)
			}
		}
	}

	return dp[0][0]
}
