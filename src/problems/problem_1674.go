// 1674. Minimum Moves to Make Array Complementary

package problems

import (
	"fmt"
)

func Problem_1674() {
	nums := []int{1, 2, 4, 3}
	limit := 4
	fmt.Println(minMoves(nums, limit))
}

func minMoves(nums []int, limit int) int {
	n := len(nums)
	diff := make([]int, 2*limit+2)
	for i := range n / 2 {
		a, b := nums[i], nums[n-1-i]
		diff[min(a, b)+1]--
		diff[a+b]--
		diff[a+b+1]++
		diff[max(a, b)+limit+1]++
	}
	moves, sum := n, n
	for i := 2; i <= 2*limit; i++ {
		sum += diff[i]
		moves = min(moves, sum)
	}
	return moves
}
