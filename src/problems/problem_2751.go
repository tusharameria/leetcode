// 2751. Robot Collisions

package problems

import (
	"fmt"
	"sort"
)

func Problem_2751() {
	positions := []int{3, 5, 2, 6}
	healths := []int{10, 10, 15, 12}
	directions := "RLRL"
	fmt.Println(survivedRobotsHealths(positions, healths, directions))
}

type RobotOne struct {
	Position  int
	Health    int
	Direction int
	Index     int
}

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	n := len(positions)
	robots := make([]RobotOne, n)
	res := make([]int, n)

	for i := 0; i < n; i++ {
		robo := RobotOne{}
		robo.Index = i
		robo.Position = positions[i]
		robo.Health = healths[i]
		currDir := directions[i]
		if currDir == 'L' {
			robo.Direction = -1
		} else {
			robo.Direction = 1
		}
		robots[i] = robo
		res[i] = robo.Health
	}

	sort.Slice(robots, func(i, j int) bool {
		return robots[i].Position < robots[j].Position
	})

	stack := []RobotOne{}

	for i := 0; i < n; i++ {
		robot := robots[i]
		if robot.Direction == 1 {
			stack = append(stack, robot)
		} else {
			currentRoboOrgIdx := robot.Index
			for len(stack) > 0 && res[currentRoboOrgIdx] > 0 {
				currentRoboHealth := res[currentRoboOrgIdx]
				stackLen := len(stack)
				stackRoboOrgIdx := stack[stackLen-1].Index
				stackRoboHealth := res[stackRoboOrgIdx]
				if stackRoboHealth > currentRoboHealth {
					res[stackRoboOrgIdx]--
					res[currentRoboOrgIdx] = 0
					if stackRoboHealth == 1 {
						stack = stack[:stackLen-1]
					}
				} else if stackRoboHealth == currentRoboHealth {
					res[stackRoboOrgIdx] = 0
					res[currentRoboOrgIdx] = 0
					stack = stack[:stackLen-1]
				} else {
					res[stackRoboOrgIdx] = 0
					res[currentRoboOrgIdx]--
					stack = stack[:stackLen-1]
				}
			}
		}
	}

	newRes := []int{}
	for _, val := range res {
		if val != 0 {
			newRes = append(newRes, val)
		}
	}

	return newRes
}
