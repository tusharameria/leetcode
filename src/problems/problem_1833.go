// 1833. Maximum Ice Cream Bars

package problems

import (
	"fmt"
	"sort"
)

func Problem_1833() {
	costs := []int{1, 6, 3, 1, 2, 5}
	coins := 20
	fmt.Println(maxIceCream(costs, coins))
}

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	count := 0
	for _, cost := range costs {
		if coins < cost {
			break
		}
		coins -= cost
		count++
	}
	return count
}
