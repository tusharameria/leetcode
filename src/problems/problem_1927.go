// 1927. Sum Game

package problems

import "fmt"

func Problem_1927() {
	num := "?6?6?000?3"
	fmt.Println(sumGame(num))
}

// Constraints:

// ---> 2 <= num.length <= 10^5
// ---> num.length is even.
// ---> num consists of only digits and '?'.

func sumGame(num string) bool {
	n := len(num)

	leftAbsSum := 0
	leftEmptyAbsCount := 0
	for i := 0; i < n/2; i++ {
		ch := num[i]
		if ch == '?' {
			leftEmptyAbsCount++
		} else {
			leftAbsSum += int(ch - '0')
		}
	}

	for i := n / 2; i < n; i++ {
		ch := num[i]
		if ch == '?' {
			leftEmptyAbsCount--
		} else {
			leftAbsSum -= int(ch - '0')
		}
	}

	fmt.Println("leftAbsSum", leftAbsSum)
	fmt.Println("leftEmptyAbsCount", leftEmptyAbsCount)

	if leftAbsSum == 0 {
		return leftEmptyAbsCount != 0
	}

	if leftEmptyAbsCount != 0 {
		if (leftAbsSum < 0) == (leftEmptyAbsCount < 0) {
			return true
		}
		if leftEmptyAbsCount%2 != 0 {
			return true
		}
		return mod(leftAbsSum) != (mod(leftEmptyAbsCount)/2)*9
	}

	return leftAbsSum != 0
}

func mod(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
