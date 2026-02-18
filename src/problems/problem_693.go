package problems

import "fmt"

// 693. Binary Number with Alternating Bits

func Problem_693() {
	n := 10
	fmt.Println("res : ", hasAlternatingBits(n))
}

func hasAlternatingBits(n int) bool {
	val := n % 2
	n /= 2
	for n > 0 {
		if n%2 == val {
			return false
		}
		val = (val + 1) % 2
		n /= 2
	}
	return true
}
