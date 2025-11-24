package problems

import "fmt"

func Problem_17() {
	digits := "2342"
	fmt.Println(letterCombinations(digits))
}

func letterCombinations(digits string) []string {
	n := len(digits)
	buff := []string{}
	for i := n - 1; i >= 0; i-- {
		buff = modify(buff, digits[i])
	}
	return buff
}

func modify(buff []string, b byte) []string {
	digitsMap := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	res := []string{}
	if len(buff) == 0 {
		for _, val := range digitsMap[b] {
			res = append(res, string([]byte{val}))
		}
	} else {
		for _, val := range digitsMap[b] {
			for i := 0; i <= len(buff)-1; i++ {
				res = append(res, string([]byte{val})+buff[i])
			}
		}
	}
	return res
}
