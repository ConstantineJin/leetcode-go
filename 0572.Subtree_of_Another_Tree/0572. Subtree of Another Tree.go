package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 104. 二叉树的最大深度
func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(getHeight(root.Left), getHeight(root.Right)) + 1
}

// 100. 相同的树
func isSameTree(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q // 必须都是 nil
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func isSubtree(root, subRoot *TreeNode) bool {
	subHeight := getHeight(subRoot)
	var dfs func(*TreeNode) (height int, found bool)
	dfs = func(node *TreeNode) (height int, found bool) {
		if node == nil {
			return 0, false
		}
		leftHeight, leftFound := dfs(node.Left)
		rightHeight, rightFound := dfs(node.Right)
		if leftFound || rightFound {
			return 0, true
		}
		nodeHeight := max(leftHeight, rightHeight) + 1
		return nodeHeight, nodeHeight == subHeight && isSameTree(node, subRoot)
	}
	_, found := dfs(root)
	return found
}
