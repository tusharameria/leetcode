// 3863. Minimum Operations to Sort a String

package problems

import "fmt"

func Problem_3863() {
	s := "salasj"
	fmt.Println(minOperations(s))
}

func minOperations(s string) int {
	sLen := len(s)
	if sLen == 1 {
		return 0
	}
	if sLen == 2 {
		if s[0] <= s[1] {
			return 0
		}
		return -1
	}

	var first byte = s[0]
	var last byte = s[sLen-1]
	nonDesc := true

	var smallest byte = 'z'
	var largest byte = 'a'

	for i := 1; i < sLen-1; i++ {
		if s[i] <= smallest {
			smallest = s[i]
		}
		if s[i] >= largest {
			largest = s[i]
		}
		if s[i] < s[i-1] {
			nonDesc = false
		}
	}

	// 1. 	f <= sm <= lg <= lt ===> if nonDesc 0 else 1
	// 1. 	f <= sm <= lt <= lg ===> 1
	// 1. 	f <= lt <= sm <= lg ===> 1
	// 1. 	sm <= f <= lg <= lt ===> 1
	// 1. 	sm <= f <= lt <= lg ===> 2
	// 1. 	sm <= lg <= f <= lt ===> 1
	// 1. 	sm <= lg <= lt <= f ===> 2
	// 1. 	sm <= lt <= f <= lg ===> 2
	// 1. 	sm <= lt <= lg <= f ===> 2
	// 1. 	lt <= f <= sm <= lg ===> 2
	// 1. 	lt <= sm <= f <= lg ===> 2
	// 1. 	lt <= sm <= lg <= f ===> 3

	if first <= smallest && largest <= last {
		if nonDesc {
			return 0
		}
		return 1
	}

	if (first <= smallest && first <= last) || (last >= largest && last >= first) {
		return 1
	}

	if first >= smallest && first <= largest && last >= smallest && last <= largest {
		return 2
	}

	if (first >= largest && first >= last && last >= smallest) || (last <= first && last <= smallest && first <= largest) {
		return 2
	}

	return 3
}
