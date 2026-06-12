/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}

	leftDept := maxDepth(root.Left)
	rightDept := maxDepth(root.Right)

	if leftDept > rightDept {
		return 1 + leftDept
	}

	return 1 + rightDept
}
