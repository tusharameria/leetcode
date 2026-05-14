// 3741. Minimum Distance Between Three Equal Elements II

package problems

import "fmt"

func Problem_3741() {
	nums := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	fmt.Println(minimumDistanceOld2(nums))
}

func minimumDistanceOld2(nums []int) int {
	numsLen := len(nums)
	distance := 2*(numsLen-1) + 1
	tupleArray := make([][]int, numsLen+1)
	for i, num := range nums {
		tupleArray[num] = append(tupleArray[num], i)
		indicesLen := len(tupleArray[num])
		if indicesLen >= 3 {
			distance = min(distance, 2*(tupleArray[num][indicesLen-1]-tupleArray[num][indicesLen-3]))
		}
	}

	if distance == 2*(numsLen-1)+1 {
		return -1
	}
	return distance
}
