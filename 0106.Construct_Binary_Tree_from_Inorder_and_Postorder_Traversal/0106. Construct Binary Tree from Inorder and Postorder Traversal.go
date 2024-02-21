package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	var val = postorder[len(inorder)-1]
	var i int
	for ; inorder[i] != val; i++ {
	}
	return &TreeNode{val, buildTree(inorder[:i], postorder[:i]), buildTree(inorder[i+1:], postorder[i:len(postorder)-1])}
}
