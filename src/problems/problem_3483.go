// 3483. Unique 3-Digit Even Numbers

package problems

import "fmt"

func Problem_3483() {
	digits := []int{2, 4, 4, 5}
	fmt.Println(totalNumbers(digits))
}

func totalNumbers(digits []int) int {
	n := len(digits)
	vis := make([]bool, 1000)
	ans := 0

	for i := 0; i < n; i++ {
		if digits[i] == 0 {
			continue
		}
		for j := 0; j < n; j++ {
			if j == i {
				continue
			}
			for k := 0; k < n; k++ {
				if k == i || k == j || digits[k]%2 != 0 {
					continue
				}
				x := digits[i]*100 + digits[j]*10 + digits[k]
				if !vis[x] {
					vis[x] = true
					ans++
				}
			}
		}
	}

	return ans
}
