// 3740. Minimum Distance Between Three Equal Elements I

package problems

func minimumDistance(nums []int) int {

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
