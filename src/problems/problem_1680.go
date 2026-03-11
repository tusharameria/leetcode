package problems

import "fmt"

func Problem_1680() {
	n := 12
	fmt.Println(concatenatedBinary(n))
}

func concatenatedBinary(n int) int {
	mod := 1000000007
	result := 0
	length := 0
	for i := 1; i <= n; i++ {
		if (i & (i - 1)) == 0 {
			length++
		}
		result = ((result << length) | i) % mod
	}
	return result
}
