// 796. Rotate String

package problems

import "fmt"

func Problem_796() {
	s := "abcdef"
	goal := "cdeabf"
	fmt.Println(rotateString(s, goal))
}

func rotateString(s string, goal string) bool {
	sLen := len(s)
	goalLen := len(goal)
	if sLen != goalLen {
		return false
	}
	if s == goal {
		return true
	}
	for i := 1; i < sLen; i++ {
		if s[i] == goal[0] {
			if s[i:] == goal[:sLen-i] && s[:i] == goal[sLen-i:] {
				return true
			}
		}
	}
	return false
}
