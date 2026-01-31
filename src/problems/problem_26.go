package problems

func Problem_26() {
	nums := []int{1, 2, 3}
	println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	lens := len(nums)
	if lens == 0 || lens == 1 {
		return lens
	}
	if lens == 2 {
		if nums[0] == nums[1] {
			return 1
		}
		return 2
	}
	left := 1
	right := 2
	for right <= len(nums)-1 {
		if nums[left] <= nums[left-1] {
			if nums[right] > nums[left-1] {
				nums[left] = nums[right]
				left++
			}
		} else {
			left++
		}
		right++
	}
	if nums[left] > nums[left-1] {
		left++
	}
	return left
}
