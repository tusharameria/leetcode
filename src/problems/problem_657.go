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
		switch move {
		case 'U':
			vertical++
		case 'D':
			vertical--
		case 'L':
			horizontal++
		case 'R':
			horizontal--
		}
	}
	return vertical == 0 && horizontal == 0
}
