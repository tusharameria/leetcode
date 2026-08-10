// 1510. Stone Game IV

package problems

import (
	"fmt"
	"math"
)

func Problem_1510() {
	n := 7
	fmt.Println(winnerSquareGame(n))
}

// Constraints:
// 1 <= n <= 1_00_000

var dp_1510 [1_00_001]bool

func init() {
	maxNum := 1_00_000
	dp_1510[1] = true

	for i := 2; i <= maxNum; i++ {
		var maxNumOfStonesSqrt int = int(math.Sqrt(float64(i)))
		possible := false
		for j := maxNumOfStonesSqrt; j >= 1; j-- {
			if !dp_1510[i-(j*j)] {
				possible = true
				break
			}
		}
		dp_1510[i] = possible
	}
}

func winnerSquareGame(n int) bool {
	fmt.Println(dp_1510[:n+1])
	return dp_1510[n]
}
