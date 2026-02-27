package problems

import "fmt"

// 1404. Number of Steps to Reduce a Number in Binary Representation to One

func Problem_1404() {
	s := "11100100101110010100101110111010111110110010"
	fmt.Println(numSteps(s))
}

func numSteps(s string) int {
	steps := 0
	right := len(s) - 1
	for right >= 0 {
		if s[right] == '0' {
			steps++
			right--
		} else {
			break
		}
	}
	left := right - 1
	for left >= 0 {
		if s[left] == '0' {
			steps += right - left + 1
			right = left
			left = right - 1
		} else {
			left--
		}
	}
	if right-left == 1 {
		return steps
	}
	return steps + right - left + 1
}
