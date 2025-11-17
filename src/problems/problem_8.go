package problems

import "fmt"

func Problem_8() {
	s := "   - 00000 5a6"
	fmt.Println(myAtoi(s))
}

func myAtoi(s string) int {
	num := 0
	started := false
	neg := false
	for i := 0; i <= len(s)-1; i++ {
		if s[i] == 32 {
			if started {
				break
			}
			continue
		}
		if s[i] >= 48 && s[i] <= 57 {
			num = (num * 10) + int(s[i]) - 48
			if num < -2147483648 || num > 2147483647 {
				if neg {
					return -2147483648
				} else {
					return 2147483647
				}
			}
			started = true
			continue
		}
		if s[i] == 43 || s[i] == 45 {
			if started {
				break
			}
			if s[i] == 45 {
				neg = true
			}
			started = true
			continue
		}
		break
	}

	if neg {
		num = -num
	}

	if num < -2147483648 || num > 2147483647 {
		if neg {
			num = -2147483648
		} else {
			num = 2147483647
		}
	}

	return num
}
