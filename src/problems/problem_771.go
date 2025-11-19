package problems

import "fmt"

func Problem_771() {
	bits := []int{1, 0, 1, 0, 0}
	fmt.Println(isOneBitCharacter(bits))
}

func isOneBitCharacter(bits []int) bool {
	if len(bits) == 0 {
		return false
	}
	bits = bits[:len(bits)-1]
	for i := 0; i <= len(bits)-1; i++ {
		if bits[i] == 1 {
			if i == len(bits)-1 {
				return false
			}
			i++
		}
	}
	return true
}

// func valid(bits []int) bool {
// 	if len(bits) == 0 {
// 		return true
// 	} else {
// 		if bits[0] == 0 {
// 			return valid(bits[1:])
// 		} else {
// 			return !(len(bits) == 1) && valid(bits[2:])
// 		}
// 	}
// }
