// 2615. Sum of Distances

package problems

import "fmt"

func Problem_2615() {
	nums := []int{1, 3, 1, 1, 2}
	fmt.Println(distance(nums))
}

func distance(nums []int) []int64 {
	numLen := len(nums)
	result := make([]int64, numLen)
	indexSumMap := make(map[int][]int64)
	// indexMapIndex := make(map[int]map[int]int)

	for i := 0; i <= numLen-1; i++ {
		indexSumMapLen := len(indexSumMap[nums[i]])
		if indexSumMapLen == 0 {
			indexSumMap[nums[i]] = append(indexSumMap[nums[i]], int64(i))
		} else {
			indexSumMap[nums[i]] = append(indexSumMap[nums[i]], indexSumMap[nums[i]][indexSumMapLen-1]+int64(i))
		}
	}

	for i := range indexSumMap {
		indexSumMapLen := len(indexSumMap[i])
		if indexSumMapLen > 1 {
			leftSum := int64(0)
			for j := 0; j <= indexSumMapLen-1; j++ {
				currentVal := indexSumMap[i][j] - leftSum
				leftSum += currentVal
				resultIndex := currentVal
				totalSum := indexSumMap[i][indexSumMapLen-1]
				currentSum := indexSumMap[i][j]
				forwardSum := totalSum - currentSum - currentVal*int64(indexSumMapLen-1-j)
				backwardSum := currentVal*int64(j+1) - currentSum
				result[resultIndex] = forwardSum + backwardSum
			}
		}
	}

	return result
}
