package main

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	switch len(preorder) {
	case 0:
		return nil
	case 1:
		return &TreeNode{Val: preorder[0]}
	default:
		var i = slices.Index(postorder, preorder[1]) + 1
		return &TreeNode{preorder[0], constructFromPrePost(preorder[1:i+1], postorder[:i]), constructFromPrePost(preorder[i+1:], postorder[i:len(postorder)-1])}
	}
}
