package problems

import "fmt"

func Problem_10() {
	s := "mississippi"
	p := "mis*is*p*."
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
	// fmt.Println(regArray)
	return stringMatch(s, regArray)
}

func stringMatch(s string, regArray []string) bool {
	if len(regArray) == 1 {
		// fmt.Println("============== Length 1 =================")
		// fmt.Println("s : ", s)
		// fmt.Println("regArray : ", regArray)
		if len(regArray[0]) == 1 {
			if len(s) != 1 {
				return false
			}
			return regArray[0][0] == '.' || string([]byte{regArray[0][0]}) == s
		} else {
			char := regArray[0][0]
			if char == '.' {
				return true
			} else {
				for i := 0; i <= len(s)-1; i++ {
					if s[i] != char {
						return false
					}
				}
				return true
			}
		}
	} else {
		// fmt.Println("============== Length != 1 =================")
		// fmt.Println("s : ", s)
		// fmt.Println("regArray : ", regArray)
		if len(regArray) == 0 {
			return len(s) == 0
		}
		if len(regArray[0]) == 1 {
			if s == "" {
				return false
			}
			return stringMatch(s[:1], regArray[:1]) && stringMatch(s[1:], regArray[1:])
		} else {
			char := regArray[0][0]
			if s == "" {
				return stringMatch(s, regArray[1:])
			} else {
				if s[0] == char || char == '.' {
					return stringMatch(s[1:], regArray) || stringMatch(s, regArray[1:])
				} else {
					return stringMatch(s, regArray[1:])
				}
			}
		}
	}
}
