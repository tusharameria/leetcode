package problems

import "fmt"

// 1461. Check If a String Contains All Binary Codes of Size K

func Problem_1461() {
	s := "00110110"
	k := 2
	fmt.Println(hasAllCodes(s, k))
}

func hasAllCodes(s string, k int) bool {
	sum := 0
	allZeroes := false
	target := PowInt(2, k)

	if len(s) < k+target-1 {
		return false
	}

	record := make([]bool, target)

	for i := 0; i < k; i++ {
		sum *= 2
		if s[i] == '1' {
			sum += 1
		}
	}

	allZeroes = (sum == 0)
	record[sum] = true
	currentSum := sum

	for i := k; i < len(s); i++ {
		if s[i-k] == '1' {
			currentSum -= target / 2
		}
		currentSum *= 2
		if s[i] == '1' {
			currentSum += 1
		}
		if !record[currentSum] {
			if currentSum == 0 {
				allZeroes = true
			}
			record[currentSum] = true
			sum += currentSum
		}
	}

	return allZeroes && sum == target*(target-1)/2
}

func hasAllCodesOld(s string, k int) bool {
	allCombMap := make(map[string]bool)

	if len(s) < k+PowInt(2, k)-1 {
		return false
	}
	var generateAllCombs func(i int, prefix string)

	generateAllCombs = func(i int, prefix string) {
		if i == 0 {
			allCombMap[prefix] = true
			return
		}
		generateAllCombs(i-1, "1"+prefix)
		generateAllCombs(i-1, "0"+prefix)
	}

	generateAllCombs(k, "")

	fmt.Println(len(s))
	for i := 0; i+k-1 <= len(s)-1; i++ {
		delete(allCombMap, s[i:i+k])
	}
	return len(allCombMap) == 0
}
