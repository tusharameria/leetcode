// 3043. Find the Length of the Longest Common Prefix

package problems

import "fmt"

func Problem_3043() {
	arr1 := []int{5655359}
	arr2 := []int{56554}
	fmt.Println(longestCommonPrefix(arr1, arr2))
}

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	maxVal := 0
	len1 := len(arr1)
	len2 := len(arr2)
	firstArrStore := make([][]int, len1)
	for i := 0; i < len1; i++ {
		val := arr1[i]
		buff := []int{}
		for val > 0 {
			buff = append(buff, val%10)
			val /= 10
		}
		firstArrStore[i] = buff
	}
	secondArrStore := make([][]int, len2)
	for i := 0; i < len2; i++ {
		val := arr2[i]
		buff := []int{}
		for val > 0 {
			buff = append(buff, val%10)
			val /= 10
		}
		secondArrStore[i] = buff
	}
	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			maxVal = max(maxVal, getCommonPrefix(firstArrStore[i], secondArrStore[j]))
		}
	}
	return maxVal
}

func getCommonPrefix(arr1, arr2 []int) int {
	num := 0

	for i, j := len(arr1)-1, len(arr2)-1; i >= 0 && j >= 0 && arr1[i] == arr2[j]; {
		num++
		i--
		j--
	}
	return num
}
