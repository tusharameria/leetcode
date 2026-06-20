// 1840. Maximum Building Height

package problems

import (
	"fmt"
	"sort"
)

func Problem_1840() {
	n := 10
	// restrictions := [][]int{{5, 3}, {2, 5}, {7, 4}, {10, 3}}
	restrictions := [][]int{{5, 3}, {2, 1}, {7, 4}, {6, 1}}
	fmt.Println(maxBuilding(n, restrictions))
}

func maxBuilding(n int, restrictions [][]int) int {
	left := 0
	for j := 0; j < len(restrictions); j++ {
		if restrictions[j][0]-1 > restrictions[j][1] {
			restrictions[left] = restrictions[j]
			left++
		}
	}
	restrictions = restrictions[:left]

	resLen := len(restrictions)
	if resLen == 0 {
		return n - 1
	}

	sort.Slice(restrictions, func(i, j int) bool {
		return restrictions[i][0] < restrictions[j][0]
	})

	// Forward Pass
	for i := 1; i < resLen; i++ {
		leftId := restrictions[i-1][0]
		leftLim := restrictions[i-1][1]
		rightId := restrictions[i][0]
		rightLim := restrictions[i][1]
		restrictions[i][1] = min(rightLim, leftLim+rightId-leftId)
	}

	// Reverse Pass
	for i := resLen - 2; i >= 0; i-- {
		leftId := restrictions[i][0]
		leftLim := restrictions[i][1]
		rightId := restrictions[i+1][0]
		rightLim := restrictions[i+1][1]
		restrictions[i][1] = min(leftLim, rightLim+rightId-leftId)
	}

	maxVal := max(
		getMaxBuildingHeight(1, 0, restrictions[0][0], restrictions[0][1]),
		getMaxBuildingHeight(restrictions[resLen-1][0], restrictions[resLen-1][1], n, restrictions[resLen-1][1]+n-restrictions[resLen-1][0]),
	)

	for i := 1; i < resLen; i++ {
		maxVal = max(maxVal, getMaxBuildingHeight(restrictions[i-1][0], restrictions[i-1][1], restrictions[i][0], restrictions[i][1]))
	}

	return maxVal
}

// id1 < id2, that's fixed
func getMaxBuildingHeight(id1, limit1, id2, limit2 int) int {
	limitDiff := abs(limit2 - limit1)
	idDiff := id2 - id1
	maxLim := max(limit1, limit2)
	if idDiff != limitDiff {
		maxLim = maxLim + (idDiff-limitDiff)/2
	}
	return maxLim
}
