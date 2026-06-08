// 2196. Create Binary Tree From Descriptions

package problems

import (
	"fmt"
)

func Problem_2196() {
	descriptions := [][]int{
		{20, 15, 1},
		{20, 17, 0},
		{50, 20, 1},
		{50, 80, 0},
		{80, 19, 1},
	}
	fmt.Println(createBinaryTree(descriptions))
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	n := len(descriptions)
	store := make(map[int]*TreeNode, 10_000)
	parents := make(map[int]int, 1_00_000)
	parent, child, isLeft := descriptions[0][0], descriptions[0][1], descriptions[0][2]
	parentNode := &TreeNode{
		Val: parent,
	}
	childNode := &TreeNode{
		Val: child,
	}
	if isLeft == 1 {
		parentNode.Left = childNode
	} else {
		parentNode.Right = childNode
	}
	root := parentNode
	if n == 1 {
		return root
	}

	store[parent] = parentNode
	store[child] = childNode
	parents[child] = parent

	lastParent := parent

	for i := 1; i < n; i++ {
		parent, child, isLeft = descriptions[i][0], descriptions[i][1], descriptions[i][2]
		parentNode, parentNodePresent := store[parent]
		childNode, childNodePresent := store[child]

		if !parentNodePresent {
			parentNode = &TreeNode{
				Val: parent,
			}
			store[parent] = parentNode
		}

		if !childNodePresent {
			childNode = &TreeNode{
				Val: child,
			}
			store[child] = childNode
		}

		if isLeft == 1 {
			parentNode.Left = childNode
		} else {
			parentNode.Right = childNode
		}
		parents[child] = parent
		lastParent = parent
	}

	root = store[lastParent]

	for {
		val := parents[lastParent]
		if val == 0 {
			break
		}
		root = store[val]
		lastParent = val
	}

	return root
}
