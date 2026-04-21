// 2078. Two Furthest Houses With Different Colors

package problems

import "fmt"

func Problem_2078() {
	colors := []int{4, 4, 4, 11, 4, 4, 11, 4, 4, 4, 4, 4}
	fmt.Println(maxDistance(colors))
}

func maxDistance(colors []int) int {
	distance := 0
	for i := 1; i <= len(colors)-1; i++ {
		for j := 0; j < i; j++ {
			if colors[i] != colors[j] {
				distance = max(distance, i-j)
			}
		}
	}
	return distance
}
