package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func distributeCoins(root *TreeNode) (ans int) {
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		d := dfs(node.Left) + dfs(node.Right) + node.Val - 1
		ans += abs(d) // 每棵子树的根节点与其父节点的边上经过的硬币数量就是子树中节点数与边数之差的绝对值
		return d
	}
	dfs(root)
	return
}
