package problems

import "fmt"

// 29. Divide Two Integers

func Problem_29() {
	dividend := 2147483646
	divisor := 2
	println(divide(dividend, divisor))
}

func divide(dividend int, divisor int) int {
	count := 0
	support := 1
	flip := false

	// Check if the result would be negative
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		flip = true
	}

	// Turn both numbers to positive
	if divisor < 0 {
		divisor = -divisor
	}
	if dividend < 0 {
		dividend = -dividend
	}

	initDivisor := divisor

	// Main Logic
	for dividend >= divisor {
		dividend -= divisor
		count += support
		divisor += divisor
		support += support
	}

	// Adjust count for overshoot
	if dividend > initDivisor {
		count += divide(dividend, initDivisor)
	}

	fmt.Printf("count before flip : %d\n", count)

	// Handle Negative Result
	if flip {
		count = -count
	}
	// Handle Leetcode Memory Limits
	if count > 2147483647 {
		count = 2147483647
	}
	if count < -2147483648 {
		count = -2147483648
	}
	return count
}

func oldDivide(dividend int, divisor int) int {
	count := 0
	flip := false
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		flip = true
	}
	if divisor < 0 {
		divisor = -divisor
	}
	if dividend < 0 {
		dividend = -dividend
	}
	for dividend >= divisor {
		dividend -= divisor
		count++
	}
	if flip {
		count = -count
	}
	if count > 2147483647 {
		count = 2147483647
	}
	if count < -2147483648 {
		count = -2147483648
	}
	return count
}
