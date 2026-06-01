// 2126. Destroying Asteroids

package problems

import (
	"fmt"
	"sort"
)

func Problem_2126() {
	mass := 10
	asteroids := []int{3, 9, 19, 5, 21}
	fmt.Println(asteroidsDestroyed(mass, asteroids))
}

func asteroidsDestroyed(mass int, asteroids []int) bool {
	sort.Ints(asteroids)
	for len(asteroids) > 0 {
		n := len(asteroids)
		if mass >= asteroids[n-1] {
			return true
		}
		if mass < asteroids[0] {
			return false
		}
		index := findBinaryLessOrEqual(asteroids, mass)
		mass += asteroids[index]
		asteroids = append(asteroids[:index], asteroids[index+1:]...)
	}
	return true
}

func findBinaryLessOrEqual(nums []int, target int) int {
	n := len(nums)
	left := 0
	right := n - 1

	for left < right {
		mid := left + (right-left+1)/2
		midVal := nums[mid]
		if midVal == target {
			return mid
		} else if midVal < target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}
