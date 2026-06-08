// 3633. Earliest Finish Time for Land and Water Rides I

package problems

import (
	"fmt"
	"math"
)

func Problem_3633() {
	landStartTime := []int{2, 8}
	landDuration := []int{4, 1}
	waterStartTime := []int{6}
	waterDuration := []int{3}
	fmt.Println(earliestFinishTime(landStartTime, landDuration, waterStartTime, waterDuration))
}

func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	landFirst := getEarliestTime(landStartTime, landDuration, 0)
	waterSecond := getEarliestTime(waterStartTime, waterDuration, landFirst)
	waterFirst := getEarliestTime(waterStartTime, waterDuration, 0)
	landSecond := getEarliestTime(landStartTime, landDuration, waterFirst)
	return min(waterSecond, landSecond)
}

func getEarliestTime(start, duration []int, freeAt int) int {
	minEnd := math.MaxInt32
	for i := 0; i < len(start); i++ {
		minEnd = min(minEnd, max(freeAt, start[i])+duration[i])
	}
	return minEnd
}
