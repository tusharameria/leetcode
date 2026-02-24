package problems

import (
	"fmt"
	"sort"
	"strings"
)

// 761. Special Binary String

func Problem_761() {
	s := "11011000"
	fmt.Println("res : ", makeLargestSpecial(s))
}

func makeLargestSpecial(s string) string {
	if len(s) <= 2 {
		return s
	}

	count := 0
	start := 0
	var parts []string

	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			count++
		} else {
			count--
		}

		if count == 0 {
			// Found one special substring
			inner := makeLargestSpecial(s[start+1 : i])
			parts = append(parts, "1"+inner+"0")
			start = i + 1
		}
	}

	sort.Slice(parts, func(i, j int) bool {
		return parts[i] > parts[j] // descending
	})

	return strings.Join(parts, "")
}
