package problems

import (
	"fmt"
	"math/bits"
)

// 401. Binary Watch

func Problem_401() {
	turnedOn := 4
	fmt.Println(readBinaryWatch(turnedOn))
}

func readBinaryWatch(turnedOn int) []string {
	res := []string{}
	for h := uint8(0); h < 12; h++ {
		for m := uint8(0); m < 60; m++ {
			if bits.OnesCount8(h)+bits.OnesCount8(m) == turnedOn {
				res = append(res, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return res
}
