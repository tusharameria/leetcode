// 2265. Count Nodes Equal to Average of Subtree

package problems

import "fmt"

func Problem_2246() {
	fmt.Println(averageOfSubtree(nil))
}

func averageOfSubtree(root *TreeNode) int {

	var res int16

	var iter func(tree *TreeNode) (count, sum int)
	iter = func(tree *TreeNode) (count, sum int) {
		sum = tree.Val
		count = 1
		if tree.Left != nil {
			leftCount, leftSum := iter(tree.Left)
			count += leftCount
			sum += leftSum
		}
		if tree.Right != nil {
			rightCount, rightSum := iter(tree.Right)
			count += rightCount
			sum += rightSum
		}
		if tree.Val == (sum / count) {
			res++
		}
		return count, sum
	}

	iter(root)

	return int(res)
}
