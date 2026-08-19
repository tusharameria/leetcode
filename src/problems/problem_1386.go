// 1386. Cinema Seat Allocation

package problems

import "fmt"

func Problem_1386() {
	n := 4
	reservedSeats := [][]int{{4, 3}, {1, 4}, {4, 6}, {1, 7}}
	fmt.Println(maxNumberOfFamilies(n, reservedSeats))
}

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	rows := make(map[int]int)

	for _, seat := range reservedSeats {
		row := seat[0]
		col := seat[1]

		if col >= 2 && col <= 9 {
			rows[row] |= 1 << col
		}
	}

	answer := 2 * (n - len(rows))

	left := (1 << 2) | (1 << 3) | (1 << 4) | (1 << 5)
	middle := (1 << 4) | (1 << 5) | (1 << 6) | (1 << 7)
	right := (1 << 6) | (1 << 7) | (1 << 8) | (1 << 9)

	for _, mask := range rows {
		canLeft := (mask & left) == 0
		canMiddle := (mask & middle) == 0
		canRight := (mask & right) == 0

		if canLeft && canRight {
			answer += 2
		} else if canLeft || canMiddle || canRight {
			answer++
		}
	}

	return answer
}
