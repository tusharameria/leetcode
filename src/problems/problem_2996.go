// 2996. Smallest Missing Integer Greater Than Sequential Prefix Sum

package problems

import "fmt"

func Problem_2996() {
	nums := []int{1, 2, 3, 2, 5}
	fmt.Println(missingInteger(nums))
}

// Constraints:

// ---> 1 <= nums.length <= 50
// ---> 1 <= nums[i] <= 50

func missingInteger(nums []int) int {
	sum := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			sum += nums[i]
		} else {
			break
		}
	}

	seen := make(map[int]bool)

	for _, num := range nums {
		seen[num] = true
	}

	answer := sum

	for seen[answer] {
		answer++
	}

	return answer
}
