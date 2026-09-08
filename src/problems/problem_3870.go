// 3870. Count Commas in Range

package problems

import "fmt"

func Problem_3870() {
	n := 1_00_000
	fmt.Println(countCommas(n))
}

func countCommas(n int) int {
	if n <= 999 {
		return 0
	}
	return n - 999
}
