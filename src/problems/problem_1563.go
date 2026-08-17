// 1563. Stone Game V

package problems

func Problem_1563()

func stoneGameV(stoneValue []int) int {
	n := len(stoneValue)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
	}
	presum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		presum[i] = presum[i-1] + stoneValue[i-1]
	}
	sum := func(i, j int) int { return presum[j+1] - presum[i] }

	for l := 2; l <= n; l++ { // l mean length
		for i := 0; i < n-l+1; i++ {
			j := i + l - 1
			score := 0
			for k := i; k < j; k++ {
				switch {
				case sum(i, k) > sum(k+1, j):
					score = max(score, memo[k+1][j]+sum(k+1, j))
				case sum(i, k) < sum(k+1, j):
					score = max(score, memo[i][k]+sum(i, k))
				default:
					tmp := max(memo[i][k]+sum(i, k), memo[k+1][j]+sum(k+1, j))
					score = max(score, tmp)
				}
			}
			memo[i][j] = score
		}
	}
	return memo[0][n-1]
}
