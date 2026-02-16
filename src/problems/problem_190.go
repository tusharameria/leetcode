package problems

import "fmt"

// 190. Reverse Bits

func Problem_190() {
	n := 43261596
	fmt.Println("res : ", reverseBits(n))
}

func reverseBits(n int) int {
	res := 0
	exp := 31
	for n > 0 {
		rem := n % 2
		res += rem * PowInt(2, exp)
		n = n / 2
		exp--
	}
	return res
}
