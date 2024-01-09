package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var leaves1, leaves2 []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			leaves1 = append(leaves1, node.Val)
			return
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root1)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			leaves2 = append(leaves2, node.Val)
			return
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root2)
	if len(leaves1) != len(leaves2) {
		return false
	}
	for i := range leaves1 {
		if leaves1[i] != leaves2[i] {
			return false
		}
	}
	return true
}
