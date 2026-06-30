// 1888. Minimum Number of Flips to Make the Binary String Alternating

package problems

import (
	"fmt"

	"github.com/tusharameria/leetcode/src/utils"
)

func Problem_1888() {
	s := utils.RandomBinaryStringGenerator(7)
	s = "01001001101"
	fmt.Println(minFlips(s))
}

func minFlips(s string) int {
	n := len(s)
	if n == 1 {
		return 0
	}
	s = s + s
	// there can only be 2 targets
	// 10101010...
	// 01010101...

	res := 0

	// Value of switches with each possibility for original number
	startWithOne := 0
	startWithZero := 0
	for i := 0; i < n; i++ {
		digit := int(s[i] - '0')
		// For case with pattern starting from 1
		// Every odd index should be one
		if i%2 == 1 {
			if digit == 1 {
				startWithOne++
			}
		} else {
			if digit == 0 {
				startWithOne++
			}
		}
	}
	startWithZero = n - startWithOne
	fmt.Println("startWithOne")
	fmt.Println(startWithOne)
	fmt.Println("startWithZero")
	fmt.Println(startWithZero)
	res = min(startWithOne, startWithZero)

	if res == 0 {
		return 0
	}

	for i := 0; i < n; i++ {
		startWithZero = startWithOne
		if n%2 == 1 {
			if s[i] == '0' {
				startWithZero--
			} else {
				startWithZero++
			}
		}
		startWithOne = n - startWithZero
		res = min(res, startWithOne, startWithZero)
	}
	return res
}
