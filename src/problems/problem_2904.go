// 2904. Shortest and Lexicographically Smallest Beautiful String

package problems

func Problem_2904() {}

// Constraints:

// ==> 1 <= s.length <= 100
// ==> 1 <= k <= s.length

func shortestBeautifulSubstring(s string, k int) string {
	answer := ""
	left := 0
	ones := 0

	for right := 0; right < len(s); right++ {
		if s[right] == '1' {
			ones++
		}

		for ones > k {
			if s[left] == '1' {
				ones--
			}
			left++
		}

		for ones == k && s[left] == '0' {
			left++
		}

		if ones == k {
			candidate := s[left : right+1]

			if answer == "" ||
				len(candidate) < len(answer) ||
				(len(candidate) == len(answer) && candidate < answer) {
				answer = candidate
			}
		}
	}

	return answer
}
