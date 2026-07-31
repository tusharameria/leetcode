// 3014. Minimum Number of Pushes to Type Word I

package problems

import "fmt"

func Problem_3014() {
	word := "qwertyuiopasdfgh"
	fmt.Println(minimumPushesOld(word))
}

func minimumPushesOld(word string) int {
	n := len(word)
	div := n / 8
	rem := n % 8
	return (4*div + rem) * (div + 1)
}
