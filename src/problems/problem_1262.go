package problems

import (
	"fmt"
	"sort"
)

func Problem_1262() {
	nums := []int{2, 6, 2, 2, 7}
	fmt.Println(maxSumDivThree(nums))
}

func maxSumDivThree(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	sum := 0
	for i := 0; i <= n-1; i++ {
		sum += nums[i]
	}
	rem := sum % 3
	if rem == 0 {
		return sum
	} else if rem == 1 {
		twoSum := 0
		twoCount := 0
		oneSum := 0
		oneCount := 0
		for i := 0; i <= n-1; i++ {
			if nums[i]%3 == 2 {
				twoSum += nums[i]
				twoCount++
			}
			if twoCount == 2 {
				break
			}
		}
		for i := 0; i <= n-1; i++ {
			if nums[i]%3 == 1 {
				oneSum += nums[i]
				oneCount++
			}
			if oneCount == 1 {
				break
			}
		}
		if twoCount != 2 && oneCount != 1 {
			return 0
		} else if twoCount == 2 && oneCount == 1 {
			if twoSum < oneSum {
				return sum - twoSum
			} else {
				return sum - oneSum
			}
		} else if twoCount == 2 {
			return sum - twoSum
		} else if oneCount == 1 {
			return sum - oneSum
		}
	} else {
		twoSum := 0
		twoCount := 0
		oneSum := 0
		oneCount := 0
		for i := 0; i <= n-1; i++ {
			if nums[i]%3 == 2 {
				twoSum += nums[i]
				twoCount++
			}
			if twoCount == 1 {
				break
			}
		}
		for i := 0; i <= n-1; i++ {
			if nums[i]%3 == 1 {
				oneSum += nums[i]
				oneCount++
			}
			if oneCount == 2 {
				break
			}
		}
		if twoCount != 1 && oneCount != 2 {
			return 0
		} else if twoCount == 1 && oneCount == 2 {
			if twoSum < oneSum {
				return sum - twoSum
			} else {
				return sum - oneSum
			}
		} else if twoCount == 1 {
			return sum - twoSum
		} else if oneCount == 2 {
			return sum - oneSum
		}
	}
	return 0
}
