// 207. Course Schedule

package problems

import "fmt"

func Problem_207() {
	numCourses := 2
	prerequisites := [][]int{{1, 0}, {0, 1}}
	fmt.Println(canFinish(numCourses, prerequisites))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	idxStore := make([][]int16, numCourses)

	n := len(prerequisites)
	// First Pass : Students can finish the courses without any prereqs
	for i := range n {
		targetCourse := prerequisites[i][0]
		preReqCourse := prerequisites[i][1]
		idxStore[targetCourse] = append(idxStore[targetCourse], int16(preReqCourse))
	}

	// 0 means not checked
	// 1 means on the path, so cycle
	// 2 means fully checked
	state := make([]int8, numCourses)

	var canComplete func(targetCourse int16) bool
	canComplete = func(targetCourse int16) bool {
		if state[targetCourse] == 2 {
			return true
		}
		if state[targetCourse] == 1 {
			return false
		}
		state[targetCourse] = 1
		for _, preReq := range idxStore[targetCourse] {
			if !canComplete(preReq) {
				return false
			}
		}
		state[targetCourse] = 2
		return true
	}

	// Second Pass : See if prereq is already done and if not, then it's prereq and so on
	for i := range n {
		// Check if prereq is already done
		preReq := prerequisites[i][1]
		if !canComplete(int16(preReq)) {
			return false
		}
	}
	return true
}
