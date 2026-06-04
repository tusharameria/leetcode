// 3751. Total Waviness of Numbers in Range I

package problems

import "fmt"

func Problem_3751() {
	num1 := 4848
	num2 := 4848
	fmt.Println(totalWaviness(num1, num2))
}

func totalWaviness(num1 int, num2 int) int {
	waviness := 0
	for i := max(100, num1); i <= num2; i++ {
		num := i
		right := num % 10
		num /= 10
		middle := num % 10
		num /= 10
		left := num % 10
		if (left < middle && right < middle) || (left > middle && right > middle) {
			waviness++
		}
		num /= 10
		for num > 0 {
			left, middle, right = num%10, left, middle
			if (left < middle && right < middle) || (left > middle && right > middle) {
				waviness++
			}
			num /= 10
		}
	}
	return waviness
}
