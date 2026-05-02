// 788. Rotated Digits

package problems

import "fmt"

func Problem_788() {
	n := 10
	fmt.Println(rotatedDigits(n))
}

func rotatedDigits(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		isGood := true
		isSame := true
		num := i
		for num > 0 {
			fmt.Printf("num : %d\n", num)
			digit := num % 10
			num /= 10
			good, same := isAGoodInteger(digit)
			if !good {
				isGood = false
				break
			}
			if !same {
				isSame = false
			}
		}
		if !isGood {
			continue
		}
		if !isSame {
			res++
		}
	}
	return res
}

func isAGoodInteger(num int) (bool, bool) {
	if num == 3 || num == 4 || num == 7 {
		return false, false
	}
	if num == 0 || num == 1 || num == 8 {
		return true, true
	}
	return true, false
}
