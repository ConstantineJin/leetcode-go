package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	var i int
	for ; inorder[i] != preorder[0]; i++ {
	}
	return &TreeNode{Val: preorder[0], Left: buildTree(preorder[1:i+1], inorder[:i]), Right: buildTree(preorder[i+1:], inorder[i+1:])}
}
