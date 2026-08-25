// 3718. Smallest Missing Multiple of K

package problems

import "fmt"

func Problem_3718() {
	nums := []int{14, 85, 81, 74, 86, 80}
	k := 14
	fmt.Println(missingMultiple(nums, k))
}

// Constraints:

// -->>> 1 <= nums.length <= 100
// -->>> 1 <= nums[i] <= 100
// -->>> 1 <= k <= 100

func missingMultiple(nums []int, k int) int {
	store := make([]bool, 101)
	for i := 0; i < len(nums); i++ {
		val := nums[i]
		store[val] = true
	}

	multiple := k

	for ; multiple <= 100; multiple += k {
		if !store[multiple] {
			return multiple
		}
	}

	return multiple
}
