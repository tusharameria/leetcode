// 3783. Mirror Distance of an Integer

package problems

import "fmt"

func Problem_3783() {
	n := 123
	fmt.Println(mirrorDistance(n))
}

func mirrorDistance(n int) int {
	return abs(n - reverseNum(n))
}
