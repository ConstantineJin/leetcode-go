package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var left, right = inorderTraversal(root.Left), inorderTraversal(root.Right)
	return append(append(left, root.Val), right...)
}
