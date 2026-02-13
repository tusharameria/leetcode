package problems

import "fmt"

// 3714. Longest Balanced Substring II

func Problem_3714() {
	nums := "bbbccb"
	fmt.Println(longestBalanced(nums))
}

func longestBalanced(s string) int {
	res := 0
	n := len(s)

	// Variables for Case I
	contCountA := 0
	contCountB := 0
	contCountC := 0

	// Variables for Case II
	prefixSumMapAB := make(map[int]int)
	prefixSumMapAB[0] = -1
	prefixSumAB := 0
	prefixSumMapBC := make(map[int]int)
	prefixSumMapBC[0] = -1
	prefixSumBC := 0
	prefixSumMapCA := make(map[int]int)
	prefixSumMapCA[0] = -1
	prefixSumCA := 0

	// Variable for Case III
	prefixSumDiffMap := make(map[int]map[int]int)
	prefixSumDiffMap[0] = make(map[int]int)
	prefixSumDiffMap[0][0] = -1
	prefixSumDiffAB := 0
	prefixSumDiffAC := 0

	for i := 0; i <= n-1; i++ {
		if s[i] == 'a' {
			// Case I : Only one charachter is present in the longest substring
			contCountA++
			contCountB = 0
			contCountC = 0
			if contCountA > res {
				res = contCountA
			}
			// Case II : Two charachters are present in the longest substring
			prefixSumAB++
			if idx, ok := prefixSumMapAB[prefixSumAB]; ok {
				if i-idx > res {
					res = i - idx
				}
			} else {
				prefixSumMapAB[prefixSumAB] = i
			}
			prefixSumCA--
			if idx, ok := prefixSumMapCA[prefixSumCA]; ok {
				if i-idx > res {
					res = i - idx
				}
			} else {
				prefixSumMapCA[prefixSumCA] = i
			}
			prefixSumMapBC = make(map[int]int)
			prefixSumMapBC[0] = i
			prefixSumBC = 0
			// Case III : All three charachters are present in the longest substring
			prefixSumDiffAB++
			prefixSumDiffAC++
			if idx, ok := prefixSumDiffMap[prefixSumDiffAB][prefixSumDiffAC]; ok {
				if i-idx > res {
					res = i - idx
				}
			} else {
				if _, ok2 := prefixSumDiffMap[prefixSumDiffAB]; !ok2 {
					prefixSumDiffMap[prefixSumDiffAB] = make(map[int]int)
				}
				prefixSumDiffMap[prefixSumDiffAB][prefixSumDiffAC] = i
			}
		} else if s[i] == 'b' {
			// Case I : Only one charachter is present in the longest substring
			contCountA = 0
			contCountB++
			contCountC = 0
			if contCountB > res {
				res = contCountB
			}
			// Case II : Two charachters are present in the longest substring
			prefixSumBC++
			fmt.Printf("prefixSumBC : %v\n", prefixSumBC)
			fmt.Printf("prev prefixSumMapBC : %v\n", prefixSumMapBC)
			if idx, ok := prefixSumMapBC[prefixSumBC]; ok {
				if i-idx > res {
					res = i - idx
				}
			} else {
				prefixSumMapBC[prefixSumBC] = i
			}
			fmt.Printf("after prefixSumMapBC : %v\n", prefixSumMapBC)
			prefixSumAB--
			if idx, ok := prefixSumMapAB[prefixSumAB]; ok {
				if i-idx > res {
					res = i - idx
				}
			} else {
				prefixSumMapAB[prefixSumAB] = i
			}
			prefixSumMapCA = make(map[int]int)
			prefixSumMapCA[0] = i
			prefixSumCA = 0
			// Case III : All three charachters are present in the longest substring
			prefixSumDiffAB--
			if idx, ok := prefixSumDiffMap[prefixSumDiffAB][prefixSumDiffAC]; ok {
				if i-idx > res {
					res = i - idx
				}
			} else {
				if _, ok2 := prefixSumDiffMap[prefixSumDiffAB]; !ok2 {
					prefixSumDiffMap[prefixSumDiffAB] = make(map[int]int)
				}
				prefixSumDiffMap[prefixSumDiffAB][prefixSumDiffAC] = i
			}
		} else {
			// Case I : Only one charachter is present in the longest substring
			contCountA = 0
			contCountB = 0
			contCountC++
			if contCountC > res {
				res = contCountC
			}
			// Case II : Two charachters are present in the longest substring
			prefixSumCA++
			if idx, ok := prefixSumMapCA[prefixSumCA]; ok {
				if i-idx > res {
					res = i - idx
				}
			} else {
				prefixSumMapCA[prefixSumCA] = i
			}
			prefixSumBC--
			fmt.Printf("prefixSumBC : %v\n", prefixSumBC)
			fmt.Printf("prev prefixSumMapBC : %v\n", prefixSumMapBC)
			if idx, ok := prefixSumMapBC[prefixSumBC]; ok {
				if i-idx > res {
					res = i - idx
				}
			} else {
				prefixSumMapBC[prefixSumBC] = i
			}
			fmt.Printf("after prefixSumMapBC : %v\n", prefixSumMapBC)
			prefixSumMapAB = make(map[int]int)
			prefixSumMapAB[0] = i
			prefixSumAB = 0
			// Case III : All three charachters are present in the longest substring
			prefixSumDiffAC--
			if idx, ok := prefixSumDiffMap[prefixSumDiffAB][prefixSumDiffAC]; ok {
				if i-idx > res {
					res = i - idx
				}
			} else {
				if _, ok2 := prefixSumDiffMap[prefixSumDiffAB]; !ok2 {
					prefixSumDiffMap[prefixSumDiffAB] = make(map[int]int)
				}
				prefixSumDiffMap[prefixSumDiffAB][prefixSumDiffAC] = i
			}
		}
	}
	fmt.Printf("last prefixSumMapBC : %v\n", prefixSumMapBC)
	fmt.Printf("last prefixSumBC : %v\n", prefixSumBC)

	return res
}
