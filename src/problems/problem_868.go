package problems

import "fmt"

// 868. Binary Gap

func Problem_868() {
	fmt.Println("res : ", binaryGap(5))
}

func binaryGap(n int) int {
	max := 0
	sum := 0
	for n > 0 {
		if n%2 != 0 {
			if n == 1 {
				return 0
			}
			n /= 2
			break
		}
		n /= 2
	}
	for n > 0 {
		if n%2 == 0 {
			sum++
		} else {
			sum++
			if sum > max {
				max = sum
			}
			sum = 0
		}
		n /= 2
	}
	return max
}
