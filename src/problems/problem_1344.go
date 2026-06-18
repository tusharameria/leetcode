// 1344. Angle Between Hands of a Clock

package problems

import "fmt"

func Problem_1344() {
	hour := 12
	minutes := 30
	fmt.Println(angleClock(hour, minutes))
}

func angleClock(hour int, minutes int) float64 {
	diff := float64((hour*30)%360) + float64((minutes*6)%360)/12 - float64((minutes*6)%360)
	if diff < 0 {
		diff = -diff
	}
	if diff > 180 {
		return 360 - diff
	}
	return diff
}
