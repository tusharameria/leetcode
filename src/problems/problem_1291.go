// 1291. Sequential Digits

package problems

import "fmt"

func Problem_1291() {
	fmt.Println(sequentialDigits(1000, 13000))
}

func sequentialDigits(low int, high int) []int {
	// All possible numbers == (8*9)/2
	store := make([]int, 36)
	// 2 digits
	store[0] = 12
	i := 1
	for ; i < 8; i++ {
		store[i] = store[i-1] + 11
	}

	// Rest of digits
	for digits := 3; digits < 10; digits++ {
		for k := 0; k < (10 - digits); k++ {
			store[i] = store[i-(11-digits)]*10 + k + digits
			i++
		}
	}

	i = 0
	for ; i < 36; i++ {
		if store[i] >= low {
			break
		}
	}
	lowIdx := i
	for ; i < 36; i++ {
		if store[i] > high {
			break
		}
	}
	highIdx := i

	return store[lowIdx:highIdx]
}
