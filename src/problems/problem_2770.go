// 2770. Maximum Number of Jumps to Reach the Last Index

package problems

import "fmt"

func Problem_2770() {
	nums := []int{1, 3, 6, 4, 1, 2}
	fmt.Println(maximumJumps(nums, 2))
}

func maximumJumps(nums []int, target int) int {
	n := len(nums)
	memory := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		found := false
		for j := i + 1; j < n; j++ {
			if memory[j] == -1 {
				continue
			}
			diff := abs(nums[j] - nums[i])
			if diff <= target {
				found = true
				if memory[i] < memory[j]+1 {
					memory[i] = memory[j] + 1
				}
			}
		}
		if !found {
			memory[i] = -1
		}
	}
	return memory[0]
}
