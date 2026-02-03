package problems

import (
	"container/heap"
	"fmt"

	ds "github.com/tusharameria/leetcode/src/dataStructures"
)

// 3013. Divide an Array Into Subarrays With Minimum Cost II

// ---------------
// Constraints:
// ---------------
// 3 <= n <= 105
// 1 <= nums[i] <= 109
// 3 <= k <= n
// k - 2 <= dist <= n - 2

func Problem_3013() {
	nums := []int{1, 5, 3, 7}
	k := 3
	dist := 1
	fmt.Printf("minimumCost : %d\n", minimumCost(nums, k, dist))
}

func minimumCost(nums []int, k int, dist int) int64 {
	cost := 0
	i := 1
	countMap := make(map[int]int)
	for i <= k-1 {
		cost += nums[i]
		countMap[nums[i]]++
		i++
	}
	tmp := make([]int, k-1)
	copy(tmp, nums[1:k])
	h := ds.MaxIntHeap(tmp)
	heap.Init(&h)

	rejectCountMap := make(map[int]int)
	rh := ds.MinIntHeap([]int{})
	heap.Init(&rh)

	lens := len(nums)
	for i <= dist+1 && i <= lens-1 {
		currentVal := nums[i]
		rejectVal := currentVal
		if currentVal < h[0] {
			if _, ok := countMap[h[0]]; ok {
				countMap[h[0]]--
				if countMap[h[0]] <= 0 {
					delete(countMap, h[0])
				}
			}
			cost -= h[0] - currentVal
			rejectVal = h[0]
			heap.Pop(&h)
			heap.Push(&h, currentVal)
			countMap[currentVal]++
		}
		heap.Push(&rh, rejectVal)
		rejectCountMap[rejectVal]++
		i++
	}

	currentCost := cost
	for i <= lens-1 {
		prevVal := nums[i-dist-1]
		currentVal := nums[i]

		if _, ok := countMap[prevVal]; ok {
			if _, ok2 := rejectCountMap[prevVal]; ok2 {
				rejectCountMap[prevVal]--
				if rejectCountMap[prevVal] <= 0 {
					delete(rejectCountMap, prevVal)
				}
				topValQ := -1
				for {
					if len(h) == 0 {
						fmt.Println("something is wrong 1")
						return -1
					}
					if _, ok := countMap[h[0]]; ok {
						topValQ = h[0]
						break
					} else {
						heap.Pop(&h)
					}
				}
				if currentVal < topValQ {
					currentCost -= topValQ - currentVal
					if currentCost < cost {
						cost = currentCost
					}
					rejectCountMap[topValQ]++
					heap.Push(&rh, topValQ)
					countMap[topValQ]--
					if countMap[topValQ] <= 0 {
						delete(countMap, topValQ)
					}
					heap.Pop(&h)
					countMap[currentVal]++
					heap.Push(&h, currentVal)
				} else {
					rejectCountMap[currentVal]++
					heap.Push(&rh, currentVal)
				}
			} else {
				countMap[prevVal]--
				if countMap[prevVal] <= 0 {
					delete(countMap, prevVal)
				}
				topValQ := -1
				for {
					if len(h) == 0 {
						fmt.Println("something is wrong 2")
						return -1
					}
					if _, ok3 := countMap[h[0]]; ok3 {
						topValQ = h[0]
						break
					} else {
						heap.Pop(&h)
					}
				}
				if currentVal < topValQ {
					countMap[currentVal]++
					heap.Push(&h, currentVal)
					currentCost -= prevVal - currentVal
					if currentCost < cost {
						cost = currentCost
					}
				} else {
					topValRQ := -1
					for len(rh) > 0 {
						if _, ok3 := rejectCountMap[rh[0]]; ok3 {
							topValRQ = rh[0]
							break
						} else {
							heap.Pop(&rh)
						}
					}
					if currentVal <= topValRQ || topValRQ == -1 {
						currentCost -= prevVal - currentVal
						if currentCost < cost {
							cost = currentCost
						}
						countMap[currentVal]++
						heap.Push(&h, currentVal)
					} else {
						currentCost -= prevVal - topValRQ
						if currentCost < cost {
							cost = currentCost
						}
						countMap[topValRQ]++
						heap.Push(&h, topValRQ)
						rejectCountMap[topValRQ]--
						if rejectCountMap[topValRQ] <= 0 {
							delete(rejectCountMap, topValRQ)
						}
						heap.Pop(&rh)
						rejectCountMap[currentVal]++
						heap.Push(&rh, currentVal)
					}
				}
			}
		} else {
			rejectCountMap[prevVal]--
			if rejectCountMap[prevVal] <= 0 {
				delete(rejectCountMap, prevVal)
			}
			topValQ := -1
			for {
				if len(h) == 0 {
					fmt.Println("something is wrong")
					return -1
				}
				if _, ok2 := countMap[h[0]]; ok2 {
					topValQ = h[0]
					break
				} else {
					heap.Pop(&h)
				}
			}
			if currentVal >= topValQ {
				rejectCountMap[currentVal]++
				heap.Push(&rh, currentVal)
			} else {
				currentCost -= topValQ - currentVal
				if currentCost < cost {
					cost = currentCost
				}
				rejectCountMap[topValQ]++
				heap.Push(&rh, topValQ)
				countMap[topValQ]--
				if countMap[topValQ] <= 0 {
					delete(countMap, topValQ)
				}
				heap.Pop(&h)
				countMap[currentVal]++
				heap.Push(&h, currentVal)
			}
		}
		i++
	}

	return int64(nums[0] + cost)
}
