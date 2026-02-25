package problems

import "fmt"

// 1022. Sum of Root To Leaf Binary Numbers

func Problem_1022() {
	fmt.Printf("res : %d\n", sumRootToLeaf(generateBinaryTree([]int{1, 0, 1, 0, 1, 0, 1})))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sumRootToLeaf(root *TreeNode) int {
	var dfs func(node *TreeNode, num int) int

	dfs = func(node *TreeNode, num int) int {
		num = (num * 2) + node.Val

		if node.Left == nil && node.Right == nil {
			return num
		}

		sum := 0
		if node.Left != nil {
			sum += dfs(node.Left, num)
		}

		if node.Right != nil {
			sum += dfs(node.Right, num)
		}
		return sum
	}

	return dfs(root, 0)
}
