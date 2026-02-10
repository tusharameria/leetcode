package problems

// 3719. Longest Balanced Subarray I

func Problem_3719() {
	nums := []int{1, 2, 3, 4, 5, 6, 6}
	println(longestBalanced(nums))
}

func longestBalanced(nums []int) int {
	maxi := 0
	for i := 0; i <= len(nums)-1; i++ {
		counts := make(map[int]int)
		counts[nums[i]]++
		distinctCounts := make(map[int]int)
		distinctCounts[nums[i]%2]++
		for j := i + 1; j <= len(nums)-1; j++ {
			if _, ok := counts[nums[j]]; !ok {
				distinctCounts[nums[j]%2]++
			}
			counts[nums[j]]++
			if distinctCounts[0] == distinctCounts[1] {
				maxi = max(maxi, j-i+1)
			}
		}
	}
	return maxi
}
