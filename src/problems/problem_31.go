package problems

//

func Problem_31() {
	nums := []int{2, 3, 4, 6, 8, 7, 5}
	nextPermutation(nums)
}

func nextPermutation(nums []int) {
	lens := len(nums)
	left := lens - 1
	right := lens - 1

	for left >= 1 {
		if nums[left] <= nums[left-1] {
			left--
		} else {
			for {
				if nums[right] <= nums[left-1] {
					right--
				} else {
					nums[left-1], nums[right] = nums[right], nums[left-1]
					break
				}
			}
			break
		}
	}

	right = lens - 1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
