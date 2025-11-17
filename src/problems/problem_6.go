package problems

import "fmt"

func Problem_6() {
	s := "PAYPALISHIRING"
	numRows := 2
	fmt.Println(convert(s, numRows))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	res := []byte{}

	for row := 0; row <= numRows-1; row++ {
		if row == 0 || row == numRows-1 {
			for i := row; i <= len(s)-1; i = i + 2*(numRows-1) {
				res = append(res, s[i])
			}
		} else {
			i := row
			for {
				if i <= len(s)-1 {
					res = append(res, s[i])
				} else {
					break
				}
				i += 2 * (numRows - row - 1)
				if i <= len(s)-1 {
					res = append(res, s[i])
				}
				i += 2 * row
			}
		}
	}
	return string(res)
}
