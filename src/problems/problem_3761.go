// 3761. Minimum Absolute Distance Between Mirror Pairs

package problems

import "fmt"

func Problem_3761() {
	nums := []int{12, 210}
	fmt.Println(minMirrorPairDistance(nums))
}

func minMirrorPairDistance(nums []int) int {
	reverseMap := make(map[int]int)
	distance := len(nums)
	for i := 0; i <= len(nums)-1; i++ {
		if v, ok := reverseMap[nums[i]]; ok {
			distance = min(distance, i-v)
			if distance == 1 {
				return 1
			}
		}
		reverseMap[reverseNum(nums[i])] = i
	}
	if distance == len(nums) {
		return -1
	}
	return distance
}

func reverseNum(num int) int {
	reversed := 0
	for num > 0 {
		reversed = reversed*10 + (num % 10)
		num /= 10
	}
	return reversed
}
