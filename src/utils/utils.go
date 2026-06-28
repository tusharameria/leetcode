package utils

import (
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
	return res
}
