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
	res := []byte{}
	n := len(s)
	left, right := 0, 0
	for right < n {
		char := s[right]
		if char >= 'a' && char <= 'z' {
			right++
		} else if char == '*' {
			j := right + 1
			for j < n && s[j] == '*' {
				j++
			}
			if left != right {
				res = append(res, []byte(s[left:right])...)
			}
			resLen := len(res)
			if j-right >= resLen {
				res = []byte{}
			} else {
				res = res[:resLen-j+right]
			}
			right, left = j, j
		} else if char == '#' {
			j := right + 1
			for j < n && s[j] == '#' {
				j++
			}
			if left != right {
				res = append(res, []byte(s[left:right])...)
			}
			for k := 0; k < j-right; k++ {
				res = append(res, res...)
			}
			right, left = j, j
		} else {
			j := right + 1
			for j < n && s[j] == '%' {
				j++
			}
			if left != right {
				res = append(res, []byte(s[left:right])...)
			}
			if (j-right)%2 == 1 {
				resLen := len(res)
				for i := 0; i < resLen/2; i++ {
					res[i], res[resLen-1-i] = res[resLen-1-i], res[i]
				}
			}
			right, left = j, j
		}
	}

	if left != right {
		res = append(res, []byte(s[left:right])...)
	}

	return string(res)
}
