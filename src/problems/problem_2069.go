// 2069. Walking Robot Simulation II

package problems

import "fmt"

func Problem_2069() {
	robot := Constructor(6, 3)
	fmt.Printf("Pos : %v, Dir : %s\n", robot.GetPos(), robot.GetDir())
	commands := []string{"step", "step", "getPos", "getDir", "step", "step", "step", "getPos", "getDir"}
	inputs := []int{2, 2, 0, 0, 2, 1, 4, 0, 0}
	for i := range commands {
		if commands[i] == "step" {
			robot.Step(inputs[i])
			fmt.Printf("Step : %d, Pos : %v, Dir : %s\n", inputs[i], robot.GetPos(), robot.GetDir())
		} else if commands[i] == "getPos" {
			robot.GetPos()
			fmt.Printf("Pos : %v\n", robot.GetPos())
		} else if commands[i] == "getDir" {
			robot.GetDir()
			fmt.Printf("Dir : %s\n", robot.GetDir())
		}
	}
}

type Robot struct {
	id          int
	moved       bool
	pos         [][2]int
	direction   []int // 0 : East, 1 : North, 2 : West, 3 : South
	toDirection map[int]string
}

func Constructor(width int, height int) Robot {
	robo := Robot{
		toDirection: map[int]string{
			0: "East",
			1: "North",
			2: "West",
			3: "South",
		},
	}

	for i := 0; i < width; i++ {
		robo.pos = append(robo.pos, [2]int{i, 0})
		robo.direction = append(robo.direction, 0)
	}
	for i := 1; i < height; i++ {
		robo.pos = append(robo.pos, [2]int{width - 1, i})
		robo.direction = append(robo.direction, 1)
	}
	for i := width - 2; i >= 0; i-- {
		robo.pos = append(robo.pos, [2]int{i, height - 1})
		robo.direction = append(robo.direction, 2)
	}
	for i := height - 2; i > 0; i-- {
		robo.pos = append(robo.pos, [2]int{0, i})
		robo.direction = append(robo.direction, 3)
	}
	robo.direction[0] = 3
	return robo
}

func (this *Robot) Step(num int) {
	this.moved = true
	this.id = (this.id + num) % len(this.pos)
}

func (this *Robot) GetPos() []int {
	return []int{this.pos[this.id][0], this.pos[this.id][1]}
}

func (this *Robot) GetDir() string {
	if !this.moved {
		return "East"
	}
	return this.toDirection[this.direction[this.id]]
}

/**
 * Your Robot object will be instantiated and called as such:
 * obj := Constructor(width, height);
 * obj.Step(num);
 * param_2 := obj.GetPos();
 * param_3 := obj.GetDir();
 */
