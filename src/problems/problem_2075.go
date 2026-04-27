// 2075. Decode the Slanted Ciphertext

package problems

import (
	"fmt"
	"strings"
	"unicode"
)

func Problem_2075() {
	encodedText := ""
	rows := 590
	fmt.Println(decodeCiphertext(encodedText, rows))
}

func decodeCiphertext(encodedText string, rows int) string {
	if rows == 1 || encodedText == "" {
		return encodedText
	}
	cols := len(encodedText) / rows
	var sb strings.Builder
	for i := 0; i <= cols-rows; i++ {
		for j := 0; j <= rows-1; j++ {
			sb.WriteByte(encodedText[i+(j*(cols+1))])
		}
	}
	for i := 0; i <= rows-2; i++ {
		for j := 0; j <= rows-2-i; j++ {
			buff := cols - rows + 1
			sb.WriteByte(encodedText[(buff+j*(cols+1))+i])
		}
	}
	return strings.TrimRightFunc(sb.String(), unicode.IsSpace)
}
