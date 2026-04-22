// 657. Robot Return to Origin

package problems

import "fmt"

func Problem_657() {
	moves := "UD"
	fmt.Println(judgeCircle(moves))
}

func judgeCircle(moves string) bool {
	vertical := 0
	horizontal := 0
	for _, move := range moves {
		if move == 'U' {
			vertical++
		} else if move == 'D' {
			vertical--
		} else if move == 'L' {
			horizontal++
		} else {
			horizontal--
		}
	}
	return vertical == 0 && horizontal == 0
}
