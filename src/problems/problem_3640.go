package problems

import "fmt"

// 3640. Trionic Array II

func Problem_3640() {
	nums := []int{-434, 332, -519, -175, 917, -316, 645}
	fmt.Printf("res : %d\n", maxSumTrionic(nums))
}

func maxSumTrionic(nums []int) int64 {
	res := -INF
	currentCost := nums[0]
	forwardCost := -INF

	x, y, z := 0, 0, 0

	for i := 1; i <= len(nums)-1; i++ {
		if nums[i] == nums[i-1] {
			x, y, z = i, i, i
			currentCost = nums[i]
			forwardCost = 0
		} else if nums[i] < nums[i-1] {
			if x == y && y == z {
				x, y, z = i, i, i
				currentCost = nums[i]
				forwardCost = 0
			} else if x == z {
				if y > z { // First Decline starting
					z = i
					currentCost += nums[i]
				} else {
					y = i - 1
					z = i
					currentCost = int(forwardCost) + nums[i]
				}
			} else { // First Decline in progress
				z = i
				currentCost += nums[i]
			}
		} else {
			// First Incline starting
			if x == y && y == z {
				y = i
				if i+1 <= len(nums)-1 && nums[i] < nums[i+1] && nums[i-1] < 0 {
					currentCost = nums[i]
				} else {
					currentCost += nums[i]
				}
			} else if x == z {
				// First Incline in Progress
				if y > z {
					y = i
					if i+1 <= len(nums)-1 && nums[i] < nums[i+1] && nums[i-1] < 0 {
						currentCost = nums[i]
					} else {
						currentCost += nums[i]
					}
				} else {
					// Second Incline in Progress
					currentCost += nums[i]
					if int64(currentCost) > res {
						res = int64(currentCost)
					}
					if i+1 <= len(nums)-1 && nums[i] < nums[i+1] {
						if nums[i] < 0 {
							forwardCost = int64(nums[i])
						} else {
							if nums[i-1] < 0 {
								forwardCost = int64(nums[i])
							} else {
								forwardCost += int64(nums[i])
							}
						}
					} else {
						forwardCost += int64(nums[i])
					}
				}
			} else {
				// Second Incline starting
				x = z
				currentCost += nums[i]
				forwardCost = int64(nums[i-1])

				if i+1 <= len(nums)-1 && nums[i] < nums[i+1] && nums[i-1] < 0 {
					forwardCost = int64(nums[i])
				} else {
					forwardCost += int64(nums[i])
				}
				if int64(currentCost) > res {
					res = int64(currentCost)
				}
			}
		}
	}

	return res
}
