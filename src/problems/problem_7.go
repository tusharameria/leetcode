package problems

import (
	"fmt"
	"math"
)

func Problem_7() {
	x := -123
	fmt.Println(reverse(x))
}

func reverse(x int) int {
	num := 0
	neg := false
	if x < 0 {
		neg = true
		x = -x
	}
	for {
		num = (num * 10) + (x % 10)
		x = x / 10
		if x == 0 {
			break
		}
	}
	if num > int(math.Pow(float64(2), float64(31))-1) {
		return 0
	}
	if neg {
		num = -num
		if num < -int(math.Pow(float64(2), float64(31))) {
			return 0
		}
	}
	return num
}
