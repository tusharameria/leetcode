// 3613. Process String with Special Operations II

// Constraints:
// 1 <= s.length <= 10^5
// s consists of only lowercase English letters and special characters '*', '#', and '%'.
// 0 <= k <= 10^15
// The length of result after processing s will not exceed 10^15.

package problems

import "fmt"

func Problem_3613() {
	// s := "uosd#eb#xz%hq#####h##%eco##tc##y#yn#fvhzyeoltlw*bjs%#cvd#tnjyt#w####rbr#ptl"
	// k := int64(2432)
	s := "qe*vkg"
	k := int64(1)
	fmt.Println(processStr(s, k))
}

func processStr(s string, k int64) byte {
	resLen := int64(0)
	lastIndexAfterZero := 0
	for i, b := range s {
		if b == '*' || b == '#' || b == '%' {
			if b == '*' {
				resLen--
				if resLen <= 0 {
					resLen = 0
					lastIndexAfterZero = i + 1
				}
			} else if b == '#' {
				resLen *= 2
			}
		} else {
			resLen++
		}
	}

	if k >= resLen {
		return '.'
	}

	s = s[lastIndexAfterZero:]

	n := len(s)
	for i := n - 1; i >= 0; i-- {
		ch := s[i]
		if ch == '*' {
			resLen++
		} else if ch == '#' {
			k = k % (resLen / 2)
			resLen /= 2
		} else if ch == '%' {
			k = resLen - 1 - k
		} else {
			resLen--
		}
		if resLen == k {
			return ch
		}
	}

	return '/'
}
