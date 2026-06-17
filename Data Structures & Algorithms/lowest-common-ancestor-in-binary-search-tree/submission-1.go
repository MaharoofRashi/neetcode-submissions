/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	node := root
    for node != nil {
		switch {
			case p.Val < node.Val && q.Val < node.Val:
				node = node.Left
			case p.Val > node.Val && q.Val > node.Val:
				node = node.Right
			default:
				return node
		}
	}
	return nil
}
