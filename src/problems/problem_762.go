package problems

import "math/bits"

func Problem_762() {}

func countPrimeSetBits(left int, right int) int {
	prime := map[int]bool{
		2: true, 3: true, 5: true, 7: true,
		11: true, 13: true, 17: true, 19: true,
	}

	count := 0

	for i := left; i <= right; i++ {
		ones := bits.OnesCount(uint(i))
		if prime[ones] {
			count++
		}
	}

	return count
}
