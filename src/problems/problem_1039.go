package problems

import "fmt"

// 1039. Minimum Score Triangulation of Polygon

func Problem_1039() {
	values := []int{1, 3, 1}
	fmt.Println("res : ", minScoreTriangulation(values))
}

func minScoreTriangulation(values []int) int {
	sides := len(values)
	if sides == 3 {
		return values[0] * values[1] * values[2]
	}
	return 0
}
