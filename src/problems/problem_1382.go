package problems

import (
	"fmt"
	"sort"
)

// 1382. Balance a Binary Search Tree

func Problem_1382() {
	root := generateBinaryTree([]int{2, 1, 3, -1, -1, -1, 4})
	printBinaryTree(balanceBST(root))
}

func balanceBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	res := []int{root.Val}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Left != nil {
			res = append(res, current.Left.Val)
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			res = append(res, current.Right.Val)
			queue = append(queue, current.Right)
		}
	}

	fmt.Println("res : ", res)
	sort.Ints(res)
	fmt.Println("res : ", res)

	return generateBinarySearchTree(res)
}
