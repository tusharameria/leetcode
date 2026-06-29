// 1689. Partitioning Into Minimum Number Of Deci-Binary Numbers

package problems

import (
	"fmt"
)

func Problem_1689() {
	n := "82734"
	fmt.Println(minPartitions(n))
}

func minPartitions(n string) int {

	var getPartitions func([]byte) int
	getPartitions = func(bArr []byte) int {
		// fmt.Printf("------- bArr : %s -------\n", string(bArr))
		numLen := len(bArr)
		// fmt.Println("numLen :", numLen)
		if numLen == 0 {
			return 0
		}
		if numLen == 1 {
			return int(bArr[0] - '0')
		}
		minDigit := 9
		maxDigit := 1
		for i := 0; i < numLen; i++ {
			digit := int(bArr[i] - '0')
			if digit != 0 {
				minDigit = min(minDigit, digit)
				maxDigit = max(maxDigit, digit)
			}
		}
		if maxDigit-minDigit <= 1 {
			return maxDigit
		}
		for i := 0; i < numLen; i++ {
			b := bArr[i]
			if digit := int(b - '0'); digit != 0 {
				bArr[i] = byte('0' + int(digit-minDigit))
			}
		}
		// fmt.Println("bArr : ", string(bArr))
		start := numLen
		for i := 0; i < numLen; i++ {
			if bArr[i] != '0' {
				start = i
				break
			}
		}
		end := start + 1
		for i := numLen - 1; i > start; i-- {
			if bArr[i] != '0' {
				end = i + 1
				break
			}
		}
		// fmt.Println("start :", start)
		// fmt.Println("end :", end)
		// time.Sleep(time.Second)
		return minDigit + getPartitions(bArr[start:end])
	}
	return getPartitions([]byte(n))
}
