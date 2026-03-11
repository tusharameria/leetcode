package problems

import "fmt"

func Problem_1680() {
	n := 12
	fmt.Println(concatenatedBinary(n))
}

func concatenatedBinary(n int) int {
	remNum := 1000000007
	rem := 1
	for i := 2; i <= n; i++ {
		rem = ((rem * PowInt(2, noBinaryDigits(i))) + i) % remNum
	}
	return rem
}

func noBinaryDigits(n int) int {
	res := 0
	for n > 0 {
		res++
		n /= 2
	}
	return res
}
