// 3661. Maximum Walls Destroyed by Robots

package problems

import (
	"fmt"
	"sort"
)

func Problem_3661() {
	robots := []int{10, 2}
	distance := []int{5, 1}
	walls := []int{5, 2, 7}
	fmt.Println(maxWalls(robots, distance, walls))
}

func maxWalls(robots []int, distance []int, walls []int) int {
	// INF := 1 << 30

	robotsToDistance := make(map[int]int)

	n := len(robots)
	for i := 0; i < n; i++ {
		robotsToDistance[robots[i]] = distance[i]
	}

	sort.Ints(robots)
	sort.Ints(walls)

	sort.Slice(robots, func(i, j int) bool {
		return robots[i] < robots[j]
	})
	sort.Slice(walls, func(i, j int) bool {
		return walls[i] < walls[j]
	})

	// fmt.Println("robots : ", robots)
	// fmt.Println("distance : ", distance)
	// fmt.Println("walls : ", walls)
	dp := make([][2]int, n)

	currentPosition := robots[0]
	leftPosition := currentPosition - robotsToDistance[currentPosition]
	dp[0][0] = findLeftMaxWalls(currentPosition, leftPosition, walls)

	rightPosition := currentPosition + robotsToDistance[currentPosition]

	if n >= 2 {
		rightPosition = min(currentPosition+robotsToDistance[currentPosition], robots[1]-1)
	}
	dp[0][1] = findRighttMaxWalls(currentPosition, rightPosition, walls)

	// fmt.Printf("Robot 0 => dp[0][0] = %d, dp[0][1] = %d\n", dp[0][0], dp[0][1])

	if n >= 2 {
		for i := 1; i < n-1; i++ {
			currentPosition = robots[i]
			// fmt.Println("=======================================")
			// fmt.Println("currentPosition : ", currentPosition)
			// fmt.Println("distance : ", robotsToDistance[currentPosition])

			// CASE I : Prev left & Current left
			leftPosition = max(currentPosition-robotsToDistance[currentPosition], robots[i-1]+1)
			maxTowardsLeftWhilePrevLeft := findLeftMaxWalls(currentPosition, leftPosition, walls)
			maxTowardsLeftWhilePrevLeft += dp[i-1][0]

			// CASE II : Prev right & Current left
			leftPosition = min(currentPosition, max(currentPosition-robotsToDistance[currentPosition], robots[i-1]+robotsToDistance[robots[i-1]]+1))
			maxTowardsLeftWhilePrevRight := findLeftMaxWalls(currentPosition, leftPosition, walls)
			maxTowardsLeftWhilePrevRight += dp[i-1][1]

			dp[i][0] = max(maxTowardsLeftWhilePrevLeft, maxTowardsLeftWhilePrevRight)

			// CASE III & IV : Current Right + max of Prev right & Prev left
			rightPosition = min(currentPosition+robotsToDistance[currentPosition], robots[i+1]-1)
			dp[i][1] = findRighttMaxWalls(currentPosition, rightPosition, walls) + max(dp[i-1][0], dp[i-1][1])
			// fmt.Printf("Robot %d => dp[%d][0] = %d, dp[%d][1] = %d\n", i, i, dp[i][0], i, dp[i][1])
		}

		// --------------------
		// LAST ROBOT
		// --------------------
		currentPosition = robots[n-1]
		// fmt.Println("=======================================")
		// fmt.Println("currentPosition : ", currentPosition)
		// fmt.Println("distance : ", robotsToDistance[currentPosition])

		// CASE I : Prev left & Current left
		leftPosition = max(currentPosition-robotsToDistance[currentPosition], robots[n-2]+1)
		maxTowardsLeftWhilePrevLeft := findLeftMaxWalls(currentPosition, leftPosition, walls)
		maxTowardsLeftWhilePrevLeft += dp[n-2][0]

		// CASE II : Prev right & Current left
		leftPosition = min(currentPosition, max(currentPosition-robotsToDistance[currentPosition], robots[n-2]+robotsToDistance[robots[n-2]]+1))
		maxTowardsLeftWhilePrevRight := findLeftMaxWalls(currentPosition, leftPosition, walls)
		maxTowardsLeftWhilePrevRight += dp[n-2][1]

		dp[n-1][0] = max(maxTowardsLeftWhilePrevLeft, maxTowardsLeftWhilePrevRight)

		// CASE III & IV : Current Right + max of Prev right & Prev left
		rightPosition = currentPosition + robotsToDistance[currentPosition]
		dp[n-1][1] = findRighttMaxWalls(currentPosition, rightPosition, walls) + max(dp[n-2][0], dp[n-2][1])

		// fmt.Printf("Robot %d => dp[%d][0] = %d, dp[%d][1] = %d\n", n-1, n-1, dp[n-1][0], n, dp[n-1][1])
	}

	return max(dp[n-1][0], dp[n-1][1])
}

func findLeftMaxWalls(startPosition, leftPosition int, walls []int) int {
	if startPosition < walls[0] || leftPosition > walls[len(walls)-1] {
		return 0
	}
	wallWithRightMostIndex := RightMostIndex(walls, startPosition)
	wallWithLeftMostIndex := LeftMostIndex(walls, leftPosition)
	return wallWithRightMostIndex - wallWithLeftMostIndex + 1
}

func findRighttMaxWalls(startPosition, rightPosition int, walls []int) int {
	if rightPosition < walls[0] || startPosition > walls[len(walls)-1] {
		return 0
	}
	wallWithRightMostIndex := RightMostIndex(walls, rightPosition)
	wallWithLeftMostIndex := LeftMostIndex(walls, startPosition)
	return wallWithRightMostIndex - wallWithLeftMostIndex + 1
}
