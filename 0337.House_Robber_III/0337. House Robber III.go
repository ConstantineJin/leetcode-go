package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	var dfs func(node *TreeNode) (int, int) // 选和不选当前节点
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		l0, l1 := dfs(node.Left)
		r0, r1 := dfs(node.Right)
		return node.Val + l1 + r1, max(l0, l1) + max(r0, r1)
	}
	m, n := dfs(root)
	return max(m, n)
}
