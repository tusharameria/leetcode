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
	indexMap := make(map[int][]int64)

	for i := 0; i <= numLen-1; i++ {
		indexMap[nums[i]] = append(indexMap[nums[i]], int64(i))
	}

	for i := range indexMap {
		indexMapLen := len(indexMap[i])
		if indexMapLen > 1 {
			leftSum := int64(0)
			totalSum := int64(0)
			for j := 0; j <= indexMapLen-1; j++ {
				totalSum += indexMap[i][j]
			}
			for j := 0; j <= indexMapLen-1; j++ {
				currentVal := indexMap[i][j]
				resultIndex := currentVal
				forwardSum := totalSum - leftSum - currentVal*int64(indexMapLen-j)
				backwardSum := currentVal*int64(j) - leftSum
				result[resultIndex] = forwardSum + backwardSum
				leftSum += currentVal
			}
		}
	}

	return result
}
