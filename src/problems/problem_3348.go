// 3348. Smallest Divisible Digit Product II

package problems

import "fmt"

func Problem_3348() {
	num := "1234"
	var t int64 = 256
	fmt.Println(smallestNumber(num, t))
}

var digitFactors = [8][][]int{
	{{0, 1}},
	{{1, 1}},
	{{0, 2}},
	{{2, 1}},
	{{0, 1}, {1, 1}},
	{{3, 1}},
	{{0, 3}},
	{{1, 2}},
}

func smallestNumber(num string, t int64) string {
	if t%10 == 0 {
		return "-1"
	}
	orgFactorCounts, valid := hasOnlySingleDigitPrimeFactors(t)
	if !valid {
		return "-1"
	}
	fmt.Println(orgFactorCounts)
	currentFactorCounts := [4]int{0, 0, 0, 0}
	copy(currentFactorCounts[:], orgFactorCounts[:])

	digitCounts := [8]int{}
	for i := 7; i >= 0; i-- {
		allPresent := true
	}
	return ""
}

func hasOnlySingleDigitPrimeFactors(n int64) ([4]int, bool) {
	nums := [4]int64{2, 3, 5, 7}
	factorCounts := [4]int{0, 0, 0, 0}
	for i, num := range nums {
		for n%num == 0 {
			n /= num
			factorCounts[i]++
		}
	}
	return factorCounts, n == 1
}
