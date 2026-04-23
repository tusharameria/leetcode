// 2069. Walking Robot Simulation II

package problems

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
	for num > 0 {
		rem := this.Direction % 2
		div := this.Direction / 2
		if rem == 0 {
			// Moving in X direction
			if div == 0 {
				// Moving in positive X direction
				if this.X+1 == this.Width {
					this.Direction = (this.Direction + 1) % 4
					continue
				}
				this.X++
				num--
			} else {
				// Moving in negative X direction
				if this.X == 0 {
					this.Direction = (this.Direction + 1) % 4
					continue
				}
				this.X--
				num--
			}
		} else {
			// Moving in Y direction
			if div == 0 {
				// Moving in positive Y direction
				if this.Y+1 == this.Height {
					this.Direction = (this.Direction + 1) % 4
					continue
				}
				this.Y++
				num--
			} else {
				// Moving in negative Y direction
				if this.Y == 0 {
					this.Direction = (this.Direction + 1) % 4
					continue
				}
				this.Y--
				num--
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
