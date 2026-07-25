// 3536. Maximum Product of Two Digits

package problems

import "fmt"

func Problem_3536() {
	n := 23
	fmt.Println(maxProduct(n))
}

func maxProduct(n int) int {
	firstMax := n % 10
	n /= 10
	secondMax := n % 10
	n /= 10
	if secondMax > firstMax {
		firstMax, secondMax = secondMax, firstMax
	}

	for n > 0 {
		dig := n % 10
		if dig >= firstMax {
			firstMax, secondMax = dig, firstMax
		} else if dig > secondMax {
			secondMax = dig
		}
		n /= 10
	}

	return firstMax * secondMax
}
