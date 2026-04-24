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
	X int
	Y int
	// 0 : East, 1 : North, 2 : West, 3 : South
	Direction int
	Width     int
	Height    int
}

func Constructor(width int, height int) Robot {
	return Robot{
		X:         0,
		Y:         0,
		Direction: 0,
		Width:     width,
		Height:    height,
	}
}

func (this *Robot) Step(num int) {
	perim := 2*(this.Width+this.Height) - 4
	if num >= perim {
		num = num % perim
	}
	if num == 0 {
		if this.X == 0 && this.Y == 0 {
			this.Direction = 3
		}
		if this.X == 0 && this.Y == this.Height-1 {
			this.Direction = 2
		}
		if this.X == this.Width-1 && this.Y == this.Height-1 {
			this.Direction = 1
		}
		if this.X == this.Width-1 && this.Y == 0 {
			this.Direction = 0
		}
		return
	}
	for num > 0 {
		rem := this.Direction % 2
		div := this.Direction / 2
		if rem == 0 {
			// Moving in X direction
			if div == 0 {
				// Moving in positive X direction
				if num > this.Width-1-this.X {
					num -= this.Width - 1 - this.X
					this.X = this.Width - 1
					this.Direction = (this.Direction + 1) % 4
					fmt.Printf("num : %d, X : %d, Direction : %d\n", num, this.X, this.Direction)
					continue
				}
				this.X += num
				return
			} else {
				// Moving in negative X direction
				if num > this.X {
					num -= this.X
					this.X = 0
					this.Direction = (this.Direction + 1) % 4
					continue
				}
				this.X -= num
				return
			}
		} else {
			// Moving in Y direction
			if div == 0 {
				// Moving in positive Y direction
				if num > this.Height-1-this.Y {
					num -= this.Height - 1 - this.Y
					this.Y = this.Height - 1
					this.Direction = (this.Direction + 1) % 4
					continue
				}
				this.Y += num
				return
			} else {
				// Moving in negative Y direction
				if num > this.Y {
					num -= this.Y
					this.Y = 0
					this.Direction = (this.Direction + 1) % 4
					continue
				}
				this.Y -= num
				return
			}
		}
	}
}

func (this *Robot) GetPos() []int {
	return []int{this.X, this.Y}
}

func (this *Robot) GetDir() string {
	if this.Direction == 0 {
		return "East"
	} else if this.Direction == 1 {
		return "North"
	} else if this.Direction == 2 {
		return "West"
	} else {
		return "South"
	}
}

/**
 * Your Robot object will be instantiated and called as such:
 * obj := Constructor(width, height);
 * obj.Step(num);
 * param_2 := obj.GetPos();
 * param_3 := obj.GetDir();
 */
