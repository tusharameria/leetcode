// 874. Walking Robot Simulation

package problems

import (
	"fmt"
	"sort"
)

func Problem_874() {
	commands := []int{6, -1, -1, 6}
	obstacles := [][]int{{0, 0}}
	fmt.Println(robotSim(commands, obstacles))
}

func robotSim(commands []int, obstacles [][]int) int {
	// Current Location of the Robot
	x, y := 0, 0
	res := 0

	// Directions in the order of North, East, South, West
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	directionIndex := 0

	// Store the obstacles in a maps for O(1) access
	obstacleXmap := make(map[int][]int)
	obstacleYmap := make(map[int][]int)

	for _, obstacle := range obstacles {
		obstacleXmap[obstacle[0]] = append(obstacleXmap[obstacle[0]], obstacle[1])
		obstacleYmap[obstacle[1]] = append(obstacleYmap[obstacle[1]], obstacle[0])
	}

	for i := range obstacleXmap {
		sort.Ints(obstacleXmap[i])
	}

	for i := range obstacleYmap {
		sort.Ints(obstacleYmap[i])
	}

	for _, command := range commands {
		if command == -1 {
			directionIndex = (directionIndex + 1) % 4
		} else if command == -2 {
			directionIndex = (directionIndex + 3) % 4
		} else {
			if len(obstacleXmap[x]) == 0 && len(obstacleYmap[y]) == 0 {
				x += directions[directionIndex][0] * command
				y += directions[directionIndex][1] * command
			} else {
				// Robot will move horizontally
				if directions[directionIndex][0] != 0 {
					if len(obstacleYmap[y]) == 0 {
						x += directions[directionIndex][0] * command
					} else {
						x += binarySearchDistanceCovered(obstacleYmap[y], x, command, directions[directionIndex][0])
					}
				} else {
					if len(obstacleXmap[x]) == 0 {
						y += directions[directionIndex][1] * command
					} else {
						y += binarySearchDistanceCovered(obstacleXmap[x], y, command, directions[directionIndex][1])
					}
				}
			}
		}
		// fmt.Printf("x : %v, y : %v\n", x, y)
		// fmt.Println("=======================")
		res = max(res, x*x+y*y)
	}

	return res
}

func binarySearchDistanceCovered(obstacles []int, current, distance, direction int) int {
	left, right := BinarySearchIntRange(obstacles, current)
	// fmt.Printf("left : %v, right : %v\n", left, right)
	// fmt.Printf("current : %v\n", current)
	// fmt.Printf("direction : %v\n", direction)
	if direction == 1 {
		if right == len(obstacles) {
			return distance
		} else {
			return min(distance, obstacles[right]-current-1)
		}
	} else {
		if left == -1 {
			return -distance
		} else {
			return max(-distance, obstacles[left]-current+1)
		}
	}
}
