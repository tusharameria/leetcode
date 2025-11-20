package problems

import "fmt"

func Problem_14() {
	strs := []string{}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	res := ""
	for i := 0; i <= len(strs[0])-1; i++ {
		for j := 0; j <= len(strs)-1; j++ {
			if i == len(strs[j]) {
				return res
			} else if strs[0][i] != strs[j][i] {
				return res
			}
		}
		res += string([]byte{strs[0][i]})
	}
	return res
}
