// 1189. Maximum Number of Balloons

package problems

import (
	"fmt"
	"slices"
)

func Problem_1189() {
	text := "loonbalxballpoon"
	fmt.Println(maxNumberOfBalloons(text))
}

func maxNumberOfBalloons(text string) int {
	single := make([]uint16, 3)
	double := make([]uint16, 2)

	for _, ch := range text {
		switch ch {
		case 'a':
			single[0]++
		case 'b':
			single[1]++
		case 'n':
			single[2]++
		case 'l':
			double[0]++
		case 'o':
			double[1]++
		}
	}
	return int(min(slices.Min(single), slices.Min(double)/2))
}
