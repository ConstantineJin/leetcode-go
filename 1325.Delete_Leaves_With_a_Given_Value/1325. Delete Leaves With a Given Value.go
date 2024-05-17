package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	left, right := removeLeafNodes(root.Left, target), removeLeafNodes(root.Right, target)
	if left == nil && right == nil && root.Val == target {
		return nil
	}
	root.Left, root.Right = left, right
	return root
}
