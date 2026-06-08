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
	store := make(map[int]*TreeNode)

	root := 0

	for i := 0; i < n; i++ {
		parent, child, isLeft := descriptions[i][0], descriptions[i][1], descriptions[i][2]
		parentNode, ok := store[parent]
		if !ok {
			parentNode = &TreeNode{
				Val: parent,
			}
			store[parent] = parentNode
			root += parent
		}
		childNode, ok := store[child]
		if !ok {
			childNode = &TreeNode{
				Val: child,
			}
			store[child] = childNode
			root += child
		}

		if isLeft == 1 {
			parentNode.Left = childNode
		} else {
			parentNode.Right = childNode
		}

		root -= child
	}

	return store[root]
}
