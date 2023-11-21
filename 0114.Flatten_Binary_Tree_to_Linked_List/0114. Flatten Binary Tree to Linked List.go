package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left) // 递归展开左右子树
	flatten(root.Right)
	root.Left, root.Right = root.Right, root.Left // 右孩子已经展开为链表
	node := root
	for ; node.Right != nil; node = node.Right {
	}
	node.Right, root.Left = root.Left, nil // 左孩子添加到右孩子的叶子节点之后，左孩子置空
}
