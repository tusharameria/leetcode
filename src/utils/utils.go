package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func PowInt(base, exp int) int {
	result := 1
	for exp > 0 {
		if exp%2 == 1 {
			result *= base
		}
		base *= base
		exp /= 2
	}
	return result
}

func RandomIntArrayGenerator(minVal, maxVal, length int) []int {
	res := make([]int, length)
	randSource := rand.NewSource(time.Now().Unix())
	r := rand.New(randSource)
	for i := range length {
		res[i] = r.Intn(maxVal-minVal+1) + minVal
	}
	fmt.Println(res)
	fmt.Println("============")
	return res
}

func RandomInt2DArrayGenerator(minVal, maxVal, row, col int) [][]int {
	res := make([][]int, row)
	for i := range res {
		res[i] = make([]int, col)
	}
	randSource := rand.NewSource(time.Now().Unix())
	r := rand.New(randSource)
	for i := range row {
		for j := range col {
			res[i][j] = r.Intn(maxVal-minVal+1) + minVal
		}
	}
	for _, row := range res {
		fmt.Println(row)
	}
	fmt.Println("============")
	return res
}

func RandomBinaryStringGenerator(len int) string {
	arr := make([]byte, len)
	randSource := rand.NewSource(time.Now().Unix())
	r := rand.New(randSource)
	for i := 0; i < len; i++ {
		val := int('0') + r.Intn(2)
		arr[i] = byte(val)
	}
	s := string(arr)
	fmt.Println(s)
	fmt.Println("============")
	return s
}
