package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) (ans int) {
	var maxDepth int
	var dfs func(root *TreeNode, currentDepth int)
	dfs = func(root *TreeNode, currentDepth int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil && currentDepth > maxDepth {
			maxDepth, ans = currentDepth, root.Val
		}
		dfs(root.Left, currentDepth+1)
		dfs(root.Right, currentDepth+1)
	}
	dfs(root, 1)
	return
}
