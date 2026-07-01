// 50. Pow(x, n)

package problems

import (
	"fmt"
)

func Problem_50() {
	x := 2.000
	n := -2
	fmt.Println(myPow(x, n))
}

func myPow(x float64, n int) float64 {
	isNeg := n < 0
	n = abs(n)
	var pow func(x float64, n int) float64
	pow = func(x float64, n int) float64 {
		if n == 1 {
			return x
		}
		isOdd := n%2 == 1
		if !isOdd {
			return pow(x*x, n/2)
		}
		return pow(x*x, n/2) * x
	}
	if isNeg {
		return 1 / pow(x, n)
	}
	return pow(x, n)
}
