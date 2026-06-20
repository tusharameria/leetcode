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
	for _, b := range s {
		switch b {
		case '*':
			if resLen > 0 {
				resLen--
			}
		case '#':
			resLen *= 2
		case '%':
		default:
			resLen++
		}
	}

	if k >= resLen {
		return '.'
	}

	for i := len(s) - 1; i >= 0; i-- {
		ch := s[i]
		switch ch {
		case '*':
			resLen++
		case '#':
			k = k % (resLen / 2)
			resLen /= 2
		case '%':
			k = resLen - 1 - k
		default:
			if k+1 == resLen {
				return ch
			}
			resLen--
		}
	}

	return '.'
}
