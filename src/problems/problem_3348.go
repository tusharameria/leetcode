// 3348. Smallest Divisible Digit Product II

package problems

import "fmt"

func Problem_3348() {
	num := "1234"
	var t int64 = 256
	fmt.Println(smallestNumber(num, t))
}

var digitPrimeExponents = [10][4]int{
	{},
	{},
	{1, 0, 0, 0},
	{0, 1, 0, 0},
	{2, 0, 0, 0},
	{0, 0, 1, 0},
	{1, 1, 0, 0},
	{0, 0, 0, 1},
	{3, 0, 0, 0},
	{0, 2, 0, 0},
}

func minimumDigitsNeeded(e [4]int) int {
	a, b := e[0], e[1]
	count := e[2] + e[3]

	count += a / 3
	a %= 3

	count += b / 2
	b %= 2

	if b == 1 {
		count++
		if a > 0 {
			a--
		}
	}

	count += (a + 1) / 2
	return count
}

func removeDigitFactors(e [4]int, digit int) [4]int {
	f := digitPrimeExponents[digit]
	for i := 0; i < 4; i++ {
		e[i] -= f[i]
		if e[i] < 0 {
			e[i] = 0
		}
	}
	return e
}

func buildSmallestSuffix(required [4]int, length int) []byte {
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		remaining := length - i - 1
		for digit := 1; digit <= 9; digit++ {
			next := removeDigitFactors(required, digit)
			if minimumDigitsNeeded(next) <= remaining {
				result[i] = byte('0' + digit)
				required = next
				break
			}
		}
	}

	return result
}

func smallestNumber(num string, t int64) string {
	var required [4]int
	primes := [4]int64{2, 3, 5, 7}

	for i, prime := range primes {
		for t%prime == 0 {
			required[i]++
			t /= prime
		}
	}
	if t != 1 {
		return "-1"
	}

	n := len(num)
	remaining := make([][4]int, n+1)
	remaining[0] = required

	firstZero := n
	for i := 0; i < n; i++ {
		digit := int(num[i] - '0')
		if digit == 0 {
			firstZero = i
			break
		}
		remaining[i+1] = removeDigitFactors(remaining[i], digit)
	}

	if firstZero == n && minimumDigitsNeeded(remaining[n]) == 0 {
		return num
	}

	lastCandidate := n - 1
	if firstZero < n {
		lastCandidate = firstZero
	}

	for i := lastCandidate; i >= 0; i-- {
		startDigit := int(num[i]-'0') + 1
		suffixLength := n - i - 1

		for digit := startDigit; digit <= 9; digit++ {
			next := removeDigitFactors(remaining[i], digit)
			if minimumDigitsNeeded(next) > suffixLength {
				continue
			}

			result := make([]byte, n)
			copy(result, num[:i])
			result[i] = byte('0' + digit)
			copy(result[i+1:], buildSmallestSuffix(next, suffixLength))
			return string(result)
		}
	}

	length := n + 1
	if needed := minimumDigitsNeeded(required); needed > length {
		length = needed
	}
	return string(buildSmallestSuffix(required, length))
}
