// 3612. Process String with Special Operations I

// Constraints:
// 1 <= s.length <= 20
// s consists of only lowercase English letters and special characters *, #, and %.

package problems

import "fmt"

func Problem_3612() {
	s := "quk**a"
	fmt.Println(processStrOld(s))
}

func processStrOld(s string) string {
	res := ""

	right := 0
	count := 1
	for right < len(s) {
		ch := s[right]
		if ch != '*' && ch != '#' && ch != '%' {
			res += string(ch)
			right++
		} else {
			if ch == '#' {
				res += res
				right++
			} else if ch == '*' {
				if right+1 < len(s) && s[right+1] == '*' {
					right++
					count++
				} else {
					if len(res) > count {
						res = res[:len(res)-count]
					} else {
						res = ""
					}
					count = 1
					right++
				}
			} else {
				r := []rune(res)
				for l, rgt := 0, len(r)-1; l < rgt; l, rgt = l+1, rgt-1 {
					r[l], r[rgt] = r[rgt], r[l]
				}
				res = string(r)
				right++
			}
		}
	}

	return res
}
