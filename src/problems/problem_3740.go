// 3740. Minimum Distance Between Three Equal Elements I

package problems

import "fmt"

func Problem_3740() {
	nums := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	fmt.Println(minimumDistanceOld(nums))
}

func minimumDistanceOld(nums []int) int {

	distance := 2*(len(nums)-1) + 1
	tupleMap := make(map[int][]int)
	for i, num := range nums {
		tupleMap[num] = append(tupleMap[num], i)
		indicesLen := len(tupleMap[num])
		if indicesLen >= 3 {
			distance = min(distance, 2*(tupleMap[num][indicesLen-1]-tupleMap[num][indicesLen-3]))
		}
	}

	if distance == 2*(len(nums)-1)+1 {
		return -1
	}
	return distance
}
