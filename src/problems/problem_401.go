package problems

import "fmt"

// 401. Binary Watch

func Problem_401() {
	turnedOn := 4
	fmt.Println(readBinaryWatch(turnedOn))
}

func readBinaryWatch(turnedOn int) []string {
	if turnedOn > 8 {
		return []string{}
	}

	if turnedOn == 0 {
		return []string{"0:00"}
	}

	res := make([]string, 0)

	for i := 0; i <= turnedOn; i++ {
		hoursLed := i
		minutesLed := turnedOn - i
		if hoursLed < 4 && minutesLed < 6 && hoursLed >= 0 && minutesLed >= 0 {

			hoursString := getHoursString(hoursLed)
			minutesString := getMinutesString(minutesLed)

			for _, h := range hoursString {
				for _, m := range minutesString {
					res = append(res, fmt.Sprintf("%s:%s", h, m))
				}
			}
		}
	}

	return res
}

func getHoursString(hoursLed int) []string {
	if hoursLed == 0 {
		return []string{"0"}
	}
	res := make([]string, 0)
	if hoursLed == 1 {
		for i := 0; i <= 3; i++ {
			res = append(res, fmt.Sprintf("%d", PowInt(2, i)))
		}
	}
	if hoursLed == 2 {
		for i := 0; i <= 3; i++ {
			for j := i + 1; j <= 3; j++ {
				val := PowInt(2, i) + PowInt(2, j)
				if val < 12 {
					res = append(res, fmt.Sprintf("%d", val))
				}
			}
		}
	}
	if hoursLed == 3 {
		res = append(res, "7")
		res = append(res, "11")
	}
	return res
}

func getMinutesString(minutesLed int) []string {
	if minutesLed == 0 {
		return []string{"00"}
	}
	res := make([]string, 0)
	if minutesLed == 1 {
		for i := 0; i <= 5; i++ {
			res = append(res, fmt.Sprintf("%02d", PowInt(2, i)))
		}
	}
	if minutesLed == 2 {
		for i := 0; i <= 5; i++ {
			for j := i + 1; j <= 5; j++ {
				val := PowInt(2, i) + PowInt(2, j)
				if val < 60 {
					res = append(res, fmt.Sprintf("%02d", val))
				}
			}
		}
	}
	if minutesLed == 3 {
		for i := 0; i <= 5; i++ {
			for j := i + 1; j <= 5; j++ {
				for k := j + 1; k <= 5; k++ {
					val := PowInt(2, i) + PowInt(2, j) + PowInt(2, k)
					if val < 60 {
						res = append(res, fmt.Sprintf("%02d", val))
					}
				}
			}
		}
	}
	if minutesLed == 4 {
		for i := 0; i <= 5; i++ {
			for j := i + 1; j <= 5; j++ {
				for k := j + 1; k <= 5; k++ {
					for l := k + 1; l <= 5; l++ {
						val := PowInt(2, i) + PowInt(2, j) + PowInt(2, k) + PowInt(2, l)
						if val < 60 {
							res = append(res, fmt.Sprintf("%02d", val))
						}
					}
				}
			}
		}
	}
	if minutesLed == 5 {
		res = append(res, "31")
		res = append(res, "47")
		res = append(res, "55")
		res = append(res, "59")
	}
	return res
}
