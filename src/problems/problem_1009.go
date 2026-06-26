package problems

import (
	"fmt"
	"math/bits"

	"github.com/tusharameria/leetcode/src/utils"
)

// 1009. Complement of Base 10 Integer

func Problem_1009() {
	n := 27
	fmt.Println(bitwiseComplement(n))
}

func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}
	return utils.PowInt(2, bits.Len(uint(n))) - n - 1
}
