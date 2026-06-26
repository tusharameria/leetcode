package segmenttrees

import (
	"fmt"

	"github.com/tusharameria/leetcode/src/utils"
)

type SegmentTreeNode struct {
	Val, Left, Right int
}

func GenerateSegmentTree() {
	nums := []int{3, 1, 2, 7}
	updateQueries := [][]int{
		{1, 4},
		{2, 6},
	}
	rangedQueries := [][]int{
		{0, 3},
		{2, 3},
		{1, 2},
		{3, 3},
	}
	fmt.Println(generateRangedSumSegmentTree(nums, updateQueries))
	fmt.Println("======= Min Queries =======")
	rangedMinSegmentTree := GenerateRangedMinSegmentTree(nums)
	fmt.Println(GetRangedMinQueryValues(rangedMinSegmentTree, rangedQueries))
	fmt.Println("======= Max Queries =======")
	rangedMaxSegmentTree := GenerateRangedMaxSegmentTree(nums)
	fmt.Println(GetRangedMaxQueryValues(rangedMaxSegmentTree, rangedQueries))
}

func GenerateRangedMinSegmentTree(nums []int) []SegmentTreeNode {
	n := len(nums)
	subTreeMaxLen := 0
	for {
		pow := utils.PowInt(2, subTreeMaxLen)
		if pow >= n {
			subTreeMaxLen = utils.PowInt(2, subTreeMaxLen+1) - 1
			break
		}
		subTreeMaxLen++
	}

	segTreeArr := make([]SegmentTreeNode, subTreeMaxLen)
	var buildTree func(i, left, right int)

	buildTree = func(i, left, right int) {
		if left == right {
			segTreeArr[i] = SegmentTreeNode{
				Val:   nums[left],
				Left:  left,
				Right: right,
			}
			return
		}
		mid := (right + left) / 2
		buildTree((2*i)+1, left, mid)
		buildTree((2*i)+2, mid+1, right)
		segTreeArr[i] = SegmentTreeNode{
			Val:   min(segTreeArr[(2*i)+1].Val, segTreeArr[(2*i)+2].Val),
			Left:  left,
			Right: right,
		}
	}

	buildTree(0, 0, n-1)
	return segTreeArr
}

func GetRangedMinQueryValues(segmentTree []SegmentTreeNode, queries [][]int) []int {
	n := len(queries)
	res := make([]int, n)
	var getRangeMin func(i, start, end int) int
	getRangeMin = func(i, start, end int) int {
		node := segmentTree[i]
		left := node.Left
		right := node.Right
		if left == start && right == end {
			return node.Val
		}
		mid := (right + left) / 2
		if end <= mid {
			return getRangeMin((2*i)+1, start, end)
		} else if start > mid {
			return getRangeMin((2*i)+2, start, end)
		} else {
			return min(getRangeMin((2*i)+1, start, mid), getRangeMin((2*i)+2, mid+1, end))
		}
	}

	for i, q := range queries {
		res[i] = getRangeMin(0, q[0], q[1])
	}
	return res
}

func GenerateRangedMaxSegmentTree(nums []int) []SegmentTreeNode {
	n := len(nums)
	subTreeMaxLen := 0
	for {
		pow := utils.PowInt(2, subTreeMaxLen)
		if pow >= n {
			subTreeMaxLen = utils.PowInt(2, subTreeMaxLen+1) - 1
			break
		}
		subTreeMaxLen++
	}

	segTreeArr := make([]SegmentTreeNode, subTreeMaxLen)
	var buildTree func(i, left, right int)

	buildTree = func(i, left, right int) {
		if left == right {
			segTreeArr[i] = SegmentTreeNode{
				Val:   nums[left],
				Left:  left,
				Right: right,
			}
			return
		}
		mid := (right + left) / 2
		buildTree((2*i)+1, left, mid)
		buildTree((2*i)+2, mid+1, right)
		segTreeArr[i] = SegmentTreeNode{
			Val:   max(segTreeArr[(2*i)+1].Val, segTreeArr[(2*i)+2].Val),
			Left:  left,
			Right: right,
		}
	}

	buildTree(0, 0, n-1)
	return segTreeArr
}

func GetRangedMaxQueryValues(segmentTree []SegmentTreeNode, queries [][]int) []int {
	n := len(queries)
	res := make([]int, n)
	var getRangeMax func(i, start, end int) int
	getRangeMax = func(i, start, end int) int {
		node := segmentTree[i]
		left := node.Left
		right := node.Right
		if left == start && right == end {
			return node.Val
		}
		mid := (right + left) / 2
		if end <= mid {
			return getRangeMax((2*i)+1, start, end)
		} else if start > mid {
			return getRangeMax((2*i)+2, start, end)
		} else {
			return max(getRangeMax((2*i)+1, start, mid), getRangeMax((2*i)+2, mid+1, end))
		}
	}

	for i, q := range queries {
		res[i] = getRangeMax(0, q[0], q[1])
	}
	return res
}

func generateRangedSumSegmentTree(nums []int, updateQueries [][]int) []SegmentTreeNode {
	n := len(nums)
	subTreeMaxLen := 0
	for {
		pow := utils.PowInt(2, subTreeMaxLen)
		if pow >= n {
			subTreeMaxLen = utils.PowInt(2, subTreeMaxLen+1) - 1
			break
		}
		subTreeMaxLen++
	}

	segTreeArr := make([]SegmentTreeNode, subTreeMaxLen)
	var buildTree func(i, left, right int)

	buildTree = func(i, left, right int) {
		if left == right {
			segTreeArr[i] = SegmentTreeNode{
				Val:   nums[left],
				Left:  left,
				Right: right,
			}
			return
		}
		mid := (right + left) / 2
		buildTree((2*i)+1, left, mid)
		buildTree((2*i)+2, mid+1, right)
		segTreeArr[i] = SegmentTreeNode{
			Val:   segTreeArr[(2*i)+1].Val + segTreeArr[(2*i)+2].Val,
			Left:  left,
			Right: right,
		}
	}

	buildTree(0, 0, n-1)
	segTreeArr = updateRangedSumSegmentTree(segTreeArr, updateQueries)
	return segTreeArr
}

func updateRangedSumSegmentTree(segmentTree []SegmentTreeNode, queries [][]int) []SegmentTreeNode {

	var updateBranch func(i, targetIdx, updatedVal int)

	updateBranch = func(i, idx, updatedValue int) {
		treeNode := segmentTree[i]
		left := treeNode.Left
		right := treeNode.Right
		if left == right {
			segmentTree[i].Val = updatedValue
			return
		}
		mid := (right + left) / 2
		newI := (2 * i) + 2
		if idx <= mid {
			newI = (2 * i) + 1
		}
		updateBranch(newI, idx, updatedValue)
		segmentTree[i].Val = segmentTree[(2*i)+1].Val + segmentTree[(2*i)+2].Val
	}

	for _, q := range queries {
		updateBranch(0, q[0], q[1])
	}
	return segmentTree
}
