// 1510. Stone Game IV

package problems

import (
	"fmt"
)

func Problem_1510() {
	arr := []int{4, 2, 7, 34, 76, 84, 986, 345}
	for _, n := range arr {
		fmt.Println(winnerSquareGame(n))
	}
}

// Constraints:
// 1 <= n <= 1_00_000

var dp_1510 [1_00_001]bool

func init() {
	dp_1510[1] = true

	for i := 1; i <= 1_00_000; i++ {
		for j := 1; j*j <= i; j++ {
			if !dp_1510[i-j*j] {
				dp_1510[i] = true
				break
			}
		}
	}
}

func winnerSquareGame(n int) bool {
	// fmt.Println(dp_1510[:n+1])
	return dp_1510[n]
}
