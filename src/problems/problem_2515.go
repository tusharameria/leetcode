// 2515. Shortest Distance to Target String in a Circular Array

package problems

import "fmt"

func Problem_2515() {
	words := []string{"hello", "i", "am", "leetcode", "hello"}
	target := "hello"
	startIndex := 1
	fmt.Println(closestTarget(words, target, startIndex))
}

func closestTarget(words []string, target string, startIndex int) int {
	distance := -1
	for i := 0; i <= len(words)-1; i++ {
		if words[i] == target {
			if distance == -1 {
				distance = min(abs(i-startIndex), len(words)-abs(i-startIndex))
			} else {
				distance = min(distance, min(abs(i-startIndex), len(words)-abs(i-startIndex)))
			}
			if distance == 0 {
				return 0
			}
		}
	}
	return distance
}
