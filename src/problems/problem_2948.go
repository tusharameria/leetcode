// 2948. Make Lexicographically Smallest Array by Swapping Elements

package problems

import (
	"fmt"
	"sort"
)

func Problem_2948() {
	nums := []int{4, 52, 38, 59, 71, 27, 31, 83, 88, 10}
	limit := 14
	fmt.Println(lexicographicallySmallestArray(nums, limit))
}

func lexicographicallySmallestArray(nums []int, limit int) []int {
	n := len(nums)

	// Step 1: Create value-index pairs
	type ValueIndex struct {
		value int
		index int
	}
	pairs := make([]ValueIndex, n)
	for i := 0; i < n; i++ {
		pairs[i] = ValueIndex{value: nums[i], index: i}
	}

	// Step 2: Sort pairs by values
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].value < pairs[j].value
	})

	// Step 3: Group elements within the threshold limit
	grouped := [][]ValueIndex{}
	group := []ValueIndex{pairs[0]}

	for i := 1; i < n; i++ {
		if pairs[i].value-pairs[i-1].value <= limit {
			group = append(group, pairs[i])
		} else {
			grouped = append(grouped, group)
			group = []ValueIndex{pairs[i]}
		}
	}
	grouped = append(grouped, group)

	// Step 4: Sort each group by index and update the original array
	for _, group := range grouped {
		indices := make([]int, len(group))
		for i, pair := range group {
			indices[i] = pair.index
		}
		sort.Ints(indices)

		// Extract values and sort them
		values := make([]int, len(group))
		for i, pair := range group {
			values[i] = pair.value
		}
		sort.Ints(values)

		// Assign sorted values back to the original indices
		for i := 0; i < len(indices); i++ {
			nums[indices[i]] = values[i]
		}
	}

	return nums
}
