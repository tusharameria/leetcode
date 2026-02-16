package problems

import "fmt"

// 67. Add Binary

func Problem_67() {
	a := "11111"
	b := "11"
	fmt.Println("res : ", addBinary(a, b))
}

func addBinary(a string, b string) string {
	carry := 0
	res := ""
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0 || carry > 0; i, j = i-1, j-1 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
		}
		if j >= 0 {
			sum += int(b[j] - '0')
		}
		res = fmt.Sprintf("%d%s", sum%2, res)
		if sum >= 2 {
			carry = 1
		} else {
			carry = 0
		}
	}

	return res
}
