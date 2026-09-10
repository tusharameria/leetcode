// 2265. Count Nodes Equal to Average of Subtree

package problems

import "fmt"

func Problem_2246() {
	fmt.Println(averageOfSubtree(nil))
}

func averageOfSubtree(root *TreeNode) int {

	var res int16

	var iter func(tree *TreeNode) (count int16, sum int32)
	iter = func(tree *TreeNode) (count int16, sum int32) {
		sum = int32(tree.Val)
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
		if int32(tree.Val) == (sum / int32(count)) {
			res++
		}
		return count, sum
	}

	iter(root)

	return int(res)
}
