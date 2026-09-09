// 3871. Count Commas in Range II

package problems

func Problem_3871() {}

var store = []int64{999, 999_999, 999_999_999, 999_999_999_999, 999_999_999_999_999, 1_000_000_000_000_000}

func countCommas(n int64) int64 {
	if n <= 999 {
		return 0
	}
	var sum int64
	for i := 1; i < 6; i++ {
		if n <= store[i] {
			sum += (n - store[i-1]) * int64(i)
			break
		}
		sum += (store[i] - store[i-1]) * int64(i)
	}
	return sum
}
