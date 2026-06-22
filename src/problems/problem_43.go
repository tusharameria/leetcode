// 43. Multiply Strings

package problems

import (
	"fmt"
)

func Problem_43() {
	num1 := "123"
	num2 := "456"
	fmt.Println(multiply(num1, num2))
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	n1 := len(num1)
	n2 := len(num2)
	res := make([]uint8, n1+n2)

	for i := n2 - 1; i >= 0; i-- {
		digit1 := uint8(num2[i] - '0')
		for j := n1 - 1; j >= 0; j-- {
			digit2 := uint8(num1[j] - '0')
			// Individual Product of digits
			prod := digit1 * digit2
			// Add the already present num on this place
			prod += res[i+j+1]
			res[i+j+1] = uint8(prod % 10)
			res[i+j] += uint8(prod / 10)
		}
	}

	idx := 0
	for i, val := range res {
		if val != 0 {
			idx = i
			break
		}
	}

	b := make([]byte, n1+n2-idx)

	for i := idx; i < n1+n2; i++ {
		b[i-idx] = byte(int('0') + int(res[i]))
	}

	return string(b)
}
