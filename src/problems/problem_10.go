package problems

import "fmt"

func Problem_10() {
	s := "mississippi"
	p := "mis*is*p*ippi"
	fmt.Println(isMatch(s, p))
}

func isMatch(s string, p string) bool {
	regArray := []string{}
	for i := len(p) - 1; i >= 0; i-- {
		if p[i] == '*' {
			regArray = append([]string{string([]byte{p[i-1], p[i]})}, regArray...)
			i--
		} else {
			regArray = append([]string{string([]byte{p[i]})}, regArray...)
		}
	}
	m := len(s)
	n := len(regArray)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	dp[m][n] = true
	for j := n - 1; j >= 0; j-- {
		dp[m][j] = len(regArray[j]) == 2 && dp[m][j+1]
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			char := regArray[j][0]
			if len(regArray[j]) == 1 {
				dp[i][j] = (char == s[i] || char == '.') && dp[i+1][j+1]
			} else {
				if char == s[i] || char == '.' {
					dp[i][j] = dp[i+1][j] || dp[i][j+1]
				} else {
					dp[i][j] = dp[i][j+1]
				}
			}
		}
	}

	return dp[0][0]
}

// func stringMatch(s string, regArray []string) bool {
// 	if len(regArray) == 1 {
// 		// fmt.Println("============== Length 1 =================")
// 		// fmt.Println("s : ", s)
// 		// fmt.Println("regArray : ", regArray)
// 		if len(regArray[0]) == 1 {
// 			if len(s) != 1 {
// 				return false
// 			}
// 			return regArray[0][0] == '.' || string([]byte{regArray[0][0]}) == s
// 		} else {
// 			char := regArray[0][0]
// 			if char == '.' {
// 				return true
// 			} else {
// 				for i := 0; i <= len(s)-1; i++ {
// 					if s[i] != char {
// 						return false
// 					}
// 				}
// 				return true
// 			}
// 		}
// 	} else {
// 		// fmt.Println("============== Length != 1 =================")
// 		// fmt.Println("s : ", s)
// 		// fmt.Println("regArray : ", regArray)
// 		if len(regArray) == 0 {
// 			return len(s) == 0
// 		}
// 		if len(regArray[0]) == 1 {
// 			if s == "" {
// 				return false
// 			}
// 			return stringMatch(s[:1], regArray[:1]) && stringMatch(s[1:], regArray[1:])
// 		} else {
// 			char := regArray[0][0]
// 			if s == "" {
// 				return stringMatch(s, regArray[1:])
// 			} else {
// 				if s[0] == char || char == '.' {
// 					return stringMatch(s[1:], regArray) || stringMatch(s, regArray[1:])
// 				} else {
// 					return stringMatch(s, regArray[1:])
// 				}
// 			}
// 		}
// 	}
// }
