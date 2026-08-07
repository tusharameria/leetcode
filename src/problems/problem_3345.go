// 3345. Smallest Divisible Digit Product I

package problems

import "fmt"

func Problem_3345() {
	n, t := 11, 3
	fmt.Println(smallestNumber3345(n, t))
}

func smallestNumber3345(n int, t int) int {
	for i := n; ; i++ {
		num := i / 10
		digit := i % 10
		prod := digit
		for {
			if num <= 0 {
				if prod%t == 0 {
					return i
				}
				break
			}
			digit = num % 10
			num /= 10
			prod *= digit
		}
	}
}
