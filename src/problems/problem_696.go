package problems

import "fmt"

// 696. Count Binary Substrings

func Problem_696() {
	s := "00110011"
	fmt.Println("res : ", countBinarySubstrings(s))
}

func countBinarySubstrings(s string) int {
	countZero := 0
	countOne := 0
	if s[0] == '0' {
		countZero = 1
	} else {
		countOne = 1
	}
	targetCount := 0
	res := 0

	for i := 1; i <= len(s)-1; i++ {
		fmt.Printf("==== i : %d ====\n", i)
		if s[i] == '0' {
			countZero++
			if s[i] != s[i-1] {
				targetCount = countOne
				countOne = 0
				res++
			} else {
				if countZero <= targetCount {
					res++
				}
			}
		} else {
			countOne++
			if s[i] != s[i-1] {
				targetCount = countZero
				countZero = 0
				res++
			} else {
				if countOne <= targetCount {
					res++
				}
			}
		}
		fmt.Printf("countZero : %d\n", countZero)
		fmt.Printf("countOne : %d\n", countOne)
		fmt.Printf("targetCount : %d\n", targetCount)
		fmt.Printf("res : %d\n", res)
	}

	return res
}
