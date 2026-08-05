// 3310. Remove Methods From Project

package problems

import "fmt"

func Problem_3310() {
	n := 3
	k := 2
	invocations := [][]int{
		{1, 0},
		{2, 0},
	}
	fmt.Println(remainingMethods(n, k, invocations))
}

func remainingMethods(n int, k int, invocations [][]int) []int {
	srcStore := make([][]int, n)
	destStore := make([][]int, n)
	for _, row := range invocations {
		first := row[0]
		second := row[1]
		srcStore[first] = append(srcStore[first], second)
		destStore[second] = append(destStore[second], first)
	}

	susMethods := make([]bool, n)

	var iterator func(src int)
	iterator = func(src int) {
		if !susMethods[src] {
			susMethods[src] = true
			for _, newSrc := range srcStore[src] {
				iterator(newSrc)
			}
		}
	}
	iterator(k)

	removeSus := true
	for i, isSus := range susMethods {
		if isSus {
			for _, val := range destStore[i] {
				if !susMethods[val] {
					removeSus = false
					break
				}
			}
		}
	}

	res := make([]int, n)
	if !removeSus {
		for i := range n {
			res[i] = i
		}
		return res
	}

	cnt := 0
	for i, isSus := range susMethods {
		if !isSus {
			res[cnt] = i
			cnt++
		}
	}
	return res[:cnt]
}
