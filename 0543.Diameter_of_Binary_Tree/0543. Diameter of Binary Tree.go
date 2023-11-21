package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) (ans int) {
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1 // 后面加上1正好是0
		}
		leftDepth, rightDepth := dfs(node.Left)+1, dfs(node.Right)+1 // 递归求左右深度
		ans = max(ans, leftDepth+rightDepth)                         // 以当前节点为根的二叉树的直径为左右子树的深度之和，酌情更新答案
		return max(leftDepth, rightDepth)                            // 返回本节点的深度
	}
	dfs(root)
	return
}
