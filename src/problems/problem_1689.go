// 1689. Partitioning Into Minimum Number Of Deci-Binary Numbers

package problems

import (
	"fmt"
)

func Problem_1689() {
	n := "82273360063400"
	fmt.Println(minPartitions(n))
}

func minPartitions(n string) int {
	maxVal := 1
	for i := 0; i < len(n); i++ {
		digit := int(n[i]) - int('0')
		if digit == 9 {
			return 9
		}
		maxVal = max(maxVal, digit)
	}
	return maxVal
}
