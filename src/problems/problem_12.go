package problems

import (
	"fmt"
)

func Problem_12() {
	num := 3749
	fmt.Println(intToRoman(num))
}

func intToRoman(num int) string {
	str := ""
	numArr := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for i := 0; i <= len(numArr)-1; i++ {
		for numArr[i] <= num {
			str += romans[i]
			num -= numArr[i]
		}
	}
	return str
}
