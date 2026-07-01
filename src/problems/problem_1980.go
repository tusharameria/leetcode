// 1980. Find Unique Binary String

package problems

import "fmt"

func Problem_1980() {
	nums := []string{
		"1001010",
		"0011000",
		"1100000",
		"0001111",
		"0111010",
		"0011101",
		"0010110",
	}
	fmt.Println(findDifferentBinaryString(nums))
}

func findDifferentBinaryString(nums []string) string {
	n := len(nums)
	bArr := make([]byte, n)

	stack := make([]int, n)
	for i := 0; i < n; i++ {
		stack[i] = i
	}

	idx := 0

	for len(stack) > 0 {
		fmt.Println(stack)
		buff := [2][]int{}
		for i := 0; i < len(stack); i++ {
			if nums[stack[i]][idx] == '0' {
				buff[0] = append(buff[0], stack[i])
			} else {
				buff[1] = append(buff[1], stack[i])
			}
		}
		if len(buff[0]) < len(buff[1]) {
			bArr[idx] = '0'
			stack = buff[0]
		} else {
			bArr[idx] = '1'
			stack = buff[1]
		}
		idx++
	}

	for i := idx; i < n; i++ {
		bArr[i] = '0'
	}

	return string(bArr)
}
