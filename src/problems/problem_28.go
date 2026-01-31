package problems

// 28. Find the Index of the First Occurrence in a String

func Problem_28() {
	haystack := "hello"
	needle := "ll"
	println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
