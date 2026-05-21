// 3043. Find the Length of the Longest Common Prefix

package problems

import "fmt"

func Problem_3043() {
	arr1 := []int{10}
	arr2 := []int{17, 11}
	fmt.Println(longestCommonPrefix(arr1, arr2))
}

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	len1 := len(arr1)
	len2 := len(arr2)

	store := make(map[int]bool, min(len1, len2)*4)
	maxVal := 0
	if len1 < len2 {
		for i := 0; i < len1; i++ {
			val := arr1[i]
			for val > 0 {
				store[val] = true
				val /= 10
			}
		}

		for i := 0; i < len2; i++ {
			val := arr2[i]
			for val > 0 {
				if store[val] {
					maxVal = max(maxVal, val)
				}
				val /= 10
			}
		}
	} else {
		for i := 0; i < len2; i++ {
			val := arr2[i]
			for val > 0 {
				store[val] = true
				val /= 10
			}
		}

		for i := 0; i < len1; i++ {
			val := arr1[i]
			for val > 0 {
				if store[val] {
					maxVal = max(maxVal, val)
				}
				val /= 10
			}
		}
	}

	i := 0
	for maxVal > 0 {
		maxVal /= 10
		i++
	}

	return i
}
