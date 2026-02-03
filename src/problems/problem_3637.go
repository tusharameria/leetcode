package problems

import "fmt"

// 3637. Trionic Array I

func Problem_3637() {
	nums := []int{1, 3, 5, 4, 2, 6}
	fmt.Printf("nums : %v\n", isTrionic(nums))
}

func isTrionic(nums []int) bool {
	if nums[1] <= nums[0] {
		return false
	}
	decreased := false
	increasedAgain := false

	for i := 2; i <= len(nums)-1; i++ {
		if nums[i] == nums[i-1] {
			return false
		} else if nums[i] < nums[i-1] {
			if increasedAgain {
				return false
			}
			decreased = true
		} else {
			if decreased {
				increasedAgain = true
			}
		}
	}

	return decreased && increasedAgain
}
