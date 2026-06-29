// 3020. Find the Maximum Number of Elements in Subset

package problems

import (
	"fmt"
	"math"
	"slices"
	"sort"
)

func Problem_3020() {
	nums := []int{1, 1, 1, 1, 1, 1}
	fmt.Println(maximumLength(nums))
}

func maximumLength(nums []int) int {
	sort.Ints(nums)
	n := len(nums)

	start := n
	for i := 0; i < n; i++ {
		if nums[i] != 1 {
			start = i
			break
		}
	}

	ones := (start - 1) / 2
	nums = nums[start:]
	n = len(nums)
	if n <= 2 {
		return ones*2 + 1
	}
	numMax := slices.Max(nums)
	if numMax < 4 {
		return 1
	}

	maxNum := math.MaxInt
	maxLen := 0

	for {
		sqrt := int(math.Sqrt(float64(maxNum)))
		if sqrt <= 1 {
			break
		}
		maxLen++
		maxNum = sqrt
	}

	seed := numMax
	checkArray := make([]int, maxLen)
	cnt := 0
	for i := 0; i < maxLen; i++ {
		sqrt := int(math.Sqrt(float64(seed)))
		if sqrt <= 1 {
			break
		}
		checkArray[i] = sqrt
		seed = sqrt
		cnt++
	}
	if ones >= cnt {
		return ones*2 + 1
	}
	checkArray = checkArray[:cnt]

	var binSearchNum func(target int) (bool, int)
	binSearchNum = func(target int) (bool, int) {
		left := 0
		right := n - 1
		for left <= right {
			mid := (right + left) / 2
			if nums[mid] == target {
				return true, mid
			}
			if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return false, 0
	}

	var checkDuplicatePresent func(int) bool
	checkDuplicatePresent = func(idx int) bool {
		present := false
		if idx-1 >= 0 {
			present = present || nums[idx-1] == nums[idx]
		}
		if idx+1 < n {
			present = present || nums[idx+1] == nums[idx]
		}
		return present
	}

	var findMaxSetLenght func(i int) int
	findMaxSetLenght = func(i int) int {
		idx := i
		dupExists := checkDuplicatePresent(idx)
		if !dupExists {
			return 0
		}
		square := nums[idx] * nums[idx]
		squarePresent, newIdx := binSearchNum(square)
		if !squarePresent {
			return 0
		}
		return 1 + findMaxSetLenght(newIdx)
	}

	// len(checkArray)*2 + 1 is max value of set we can have
	squared := ones
	seed = nums[0] * nums[0]
	for m := 0; m < squared; m++ {
		seedSuare := seed * seed
		if seedSuare > numMax {
			return (squared * 2) + 1
		}
		seed = seedSuare
	}

	squared = max(squared, findMaxSetLenght(0))

	for i := 1; i < n; i++ {
		numVal := nums[i]
		if squared == cnt {
			return (squared * 2) + 1
		}
		if numVal > checkArray[squared] {
			return (squared * 2) + 1
		}
		if nums[i] == nums[i-1] {
			continue
		}
		seed := numVal
		for m := 0; m < squared+1; m++ {
			seedSquare := seed * seed
			if seedSquare > numMax {
				return (squared * 2) + 1
			}
			seed = seedSquare
		}

		squared = max(squared, findMaxSetLenght(i))
	}

	return (squared * 2) + 1
}
