package problems

// 1653. Minimum Deletions to Make String Balanced

func Problem_1653() {}

func minimumDeletions(s string) int {
	bCount := 0
	ans := 0

	for _, ch := range s {
		if ch == 'b' {
			bCount++
		} else { // ch == 'a'
			ans = min(ans+1, bCount)
		}
	}
	return ans
}
