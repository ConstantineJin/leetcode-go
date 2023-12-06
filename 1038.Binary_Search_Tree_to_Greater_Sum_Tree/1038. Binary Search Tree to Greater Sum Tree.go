package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pre int

func traversal(node *TreeNode) {
	if node == nil {
		return
	} else {
		traversal(node.Right)
		node.Val += pre
		pre = node.Val
		traversal(node.Left)
	}
}

func bstToGst(root *TreeNode) *TreeNode {
	pre = 0
	traversal(root)
	return root
}
