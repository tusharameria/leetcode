// 1358. Number of Substrings Containing All Three Characters

package problems

import "fmt"

func Problem_1358() {
	s := "abcabc"
	fmt.Println(numberOfSubstrings(s))
}

func numberOfSubstrings(s string) int {
	n := len(s)
	aIdxStore := make([]int, n)
	bIdxStore := make([]int, n)
	cIdxStore := make([]int, n)

	aCount, bCount, cCount := 0, 0, 0

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a':
			aIdxStore[aCount] = i
			aCount++
		case 'b':
			bIdxStore[bCount] = i
			bCount++
		default:
			cIdxStore[cCount] = i
			cCount++
		}
	}
	aIdxStore = aIdxStore[:aCount]
	bIdxStore = bIdxStore[:bCount]
	cIdxStore = cIdxStore[:cCount]
	aIdx, bIdx, cIdx := 0, 0, 0

	res := 0
	for {
		if aIdx == aCount || bIdx == bCount || cIdx == cCount {
			return res
		}
		aIdxVal, bIdxVal, cIdxVal := aIdxStore[aIdx], bIdxStore[bIdx], cIdxStore[cIdx]
		res += n - max(aIdxVal, bIdxVal, cIdxVal)
		if aIdxVal < bIdxVal && aIdxVal < cIdxVal {
			aIdx++
		} else if bIdxVal < cIdxVal {
			bIdx++
		} else {
			cIdx++
		}
	}
}
