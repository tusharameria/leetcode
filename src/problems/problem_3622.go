// 3622. Check Divisibility by Digit Sum and Product

package problems

import "fmt"

func Problem_3622() {
	n := 67
	fmt.Println(checkDivisibility(n))
}

func checkDivisibility(n int) bool {
	num := n
	rem := n % 10
	digSum, digProd := rem, rem
	n /= 10

	for n > 0 {
		rem = n % 10
		digSum += rem
		digProd *= rem
		n /= 10
	}

	sum := digSum + digProd

	if sum == 0 {
		return false
	}

	return num%sum == 0
}
