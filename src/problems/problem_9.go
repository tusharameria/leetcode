package problems

import (
	"fmt"
)

func Problem_9() {
	x := 10
	fmt.Println(isPalindrome(x))
}

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	rev := 0
	for x > rev {
		rev = (rev * 10) + x%10
		x = x / 10
	}
	return x == rev || x == rev/10
}
