// 2833. Furthest Point From Origin

package problems

import "fmt"

func Problem_2833() {
	moves := "_R__LL_"
	fmt.Println(furthestDistanceFromOrigin(moves))
}

func furthestDistanceFromOrigin(moves string) int {
	x := 0
	score := 0
	for i := 0; i < len(moves); i++ {
		if moves[i] == 'R' {
			x--
		} else if moves[i] == 'L' {
			x++
		} else {
			score++
		}
	}
	if x < 0 {
		return -x + score
	}
	return x + score
}
