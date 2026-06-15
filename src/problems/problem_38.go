// 38. Count and Say

package problems

import (
	"fmt"
	"strconv"
)

func Problem_38() {
	n := 3
	fmt.Println(countAndSay(n))
}

func countAndSay(n int) string {
	s := []byte{'1'}
	for i := 1; i < n; i++ {
		count := 1
		char := s[0]
		res := []byte{}
		for j := 1; j < len(s); j++ {
			if s[j] == char {
				count++
			} else {
				res = append(res, append([]byte(strconv.Itoa(count)), []byte{char}...)...)
				char = s[j]
				count = 1
			}
		}
		res = append(res, append([]byte(strconv.Itoa(count)), []byte{char}...)...)
		s = res
		res = []byte{}
	}
	return string(s)
}
