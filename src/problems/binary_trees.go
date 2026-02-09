package problems

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateBinaryTree(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	root := &TreeNode{Val: arr[0]}
	queue := []*TreeNode{root}
	i := 1
	for i < len(arr) {
		current := queue[0]
		queue = queue[1:]

		if arr[i] != -1 {
			current.Left = &TreeNode{Val: arr[i]}
			queue = append(queue, current.Left)
		}
		i++

		if i < len(arr) && arr[i] != -1 {
			current.Right = &TreeNode{Val: arr[i]}
			queue = append(queue, current.Right)
		}
		i++
	}
	return root
}

func generateBinarySearchTree(left, right int, arr []int) *TreeNode {
	if left < 0 || right >= len(arr) || left > right {
		return nil
	}
	mid := left + (right-left)/2
	root := &TreeNode{Val: arr[mid]}
	root.Left = generateBinarySearchTree(left, mid-1, arr)
	root.Right = generateBinarySearchTree(mid+1, right, arr)
	return root
}

func printBinaryTree(root *TreeNode) {
	if root == nil {
		fmt.Println("Empty tree")
	}
	queue := []*TreeNode{root}
	res := []int{root.Val}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.Left != nil {
			res = append(res, current.Left.Val)
			queue = append(queue, current.Left)
		} else {
			res = append(res, -1)
		}
		if current.Right != nil {
			res = append(res, current.Right.Val)
			queue = append(queue, current.Right)
		} else {
			res = append(res, -1)
		}
	}

	fmt.Printf("res %v\n", res)
}
