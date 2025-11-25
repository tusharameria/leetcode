package problems

import "fmt"

func Problem_1015() {
	k := 9
	fmt.Println(smallestRepunitDivByK(k))
}

func smallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}
	rem := 0

	for length := 1; length <= k; length++ {
		rem = (rem*10 + 1) % k
		if rem == 0 {
			return length
		}
	}
	return -1
}
