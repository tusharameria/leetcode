// 3754. Concatenate Non-Zero Digits and Multiply by Sum I

package problems

import (
	"fmt"
	"math"
)

func Problem_3754() {
	n := 51
	fmt.Println(sumAndMultiply(n))
}

func sumAndMultiply(n int) int64 {
	new := 0
	sum := 0
	powerOfTen := 0

	for n > 0 {
		rem := n % 10
		sum += rem
		n /= 10
		if rem == 0 {
			continue
		}
		new += int(math.Pow10(powerOfTen)) * rem
		powerOfTen++
	}

	return int64(new * sum)
}
