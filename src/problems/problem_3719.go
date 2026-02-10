package problems

import "fmt"

// 3719. Longest Balanced Subarray I

func Problem_3719() {
	nums := []int{1, 2, 3, 2, 5}
	fmt.Println(longestBalanced(nums))
}

func longestBalanced(nums []int) int {
	maxi := 0
	for i := 0; i <= len(nums)-1-maxi; i++ {
		evenSet := make(map[int]bool)
		oddSet := make(map[int]bool)
		if nums[i]%2 == 0 {
			evenSet[nums[i]] = true
		} else {
			oddSet[nums[i]] = true
		}
		for j := i + 1; j <= len(nums)-1; j++ {
			if nums[j]%2 == 0 {
				evenSet[nums[j]] = true
			} else {
				oddSet[nums[j]] = true
			}
			if len(evenSet) == len(oddSet) {
				maxi = max(maxi, j-i+1)
			}
		}
	}
	return maxi
}
