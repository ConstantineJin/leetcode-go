package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodes(root *TreeNode, toDelete []int) (ans []*TreeNode) {
	del := make(map[int]bool)
	for _, v := range toDelete {
		del[v] = true
	}
	var dfs func(node *TreeNode, fatherIsDeleted bool)
	dfs = func(node *TreeNode, fatherIsDeleted bool) {
		if node == nil {
			return
		}
		dfs(node.Left, del[node.Val])
		dfs(node.Right, del[node.Val])
		if node.Left != nil && del[node.Left.Val] {
			node.Left = nil
		}
		if node.Right != nil && del[node.Right.Val] {
			node.Right = nil
		}
		if fatherIsDeleted && !del[node.Val] {
			ans = append(ans, node)
		}
	}
	dfs(root, true)
	return
}
