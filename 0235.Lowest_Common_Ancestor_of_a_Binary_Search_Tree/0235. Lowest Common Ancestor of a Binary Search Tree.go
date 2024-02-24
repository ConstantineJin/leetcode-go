package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) (ans *TreeNode) {
	if p.Val > q.Val {
		p, q = q, p
	}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Val < p.Val {
			dfs(node.Right)
		} else if node.Val > q.Val {
			dfs(node.Left)
		} else {
			ans = node
			return
		}
	}
	dfs(root)
	return
}
