package problems

// 1382. Balance a Binary Search Tree

func Problem_1382() {
	root := generateBinaryTree([]int{2, 1, 3, -1, -1, -1, 4})
	printBinaryTree(balanceBST(root))
}

func balanceBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	res := []*TreeNode{}
	inOrderNode(root, &res)

	return generateBinarySearchTreeNode(res)
}

func inOrderNode(root *TreeNode, nodes *[]*TreeNode) {
	if root == nil {
		return
	}
	inOrderNode(root.Left, nodes)
	*nodes = append(*nodes, root)
	inOrderNode(root.Right, nodes)
}

func generateBinarySearchTreeNode(nodes []*TreeNode) *TreeNode {
	if len(nodes) == 0 {
		return nil
	}
	mid := len(nodes) / 2
	root := nodes[mid]
	root.Left = generateBinarySearchTreeNode(nodes[:mid])
	root.Right = generateBinarySearchTreeNode(nodes[mid+1:])
	return root
}
