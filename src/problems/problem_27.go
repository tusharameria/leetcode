package problems

// 27. Remove Element

func Problem_27() {
	nums := []int{3, 2, 2, 3}
	val := 3
	println(removeElement(nums, val))
}

func removeElement(nums []int, val int) int {
	lens := len(nums)
	if lens == 0 {
		return 0
	}
	left := 0
	right := 0
	for right <= lens-1 {
		if nums[right] != val {
			if nums[left] == val {
				nums[left], nums[right] = nums[right], nums[left]
			}
			left++
		}
		right++
	}
	return left
}
