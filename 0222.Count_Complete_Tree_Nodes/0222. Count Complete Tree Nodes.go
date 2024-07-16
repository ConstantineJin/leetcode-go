package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var leftHeight, rightHeight int
	node := root
	for node.Left != nil {
		node = node.Left
		leftHeight++
	}
	node = root
	for node.Right != nil {
		node = node.Right
		rightHeight++
	}
	if leftHeight == rightHeight { // 当前子树是一棵满二叉树
		return (2 << leftHeight) - 1
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}
