package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) (ans int) {
	var dfs func(node *TreeNode, mx, mn int) // mx和mn为当前自顶向下路径的最值
	dfs = func(node *TreeNode, mx, mn int) {
		if node == nil {
			ans = max(ans, mx-mn)
			return
		}
		mx, mn = max(mx, node.Val), min(mn, node.Val)
		dfs(node.Left, mx, mn)
		dfs(node.Right, mx, mn)
	}
	dfs(root, root.Val, root.Val)
	return ans
}
