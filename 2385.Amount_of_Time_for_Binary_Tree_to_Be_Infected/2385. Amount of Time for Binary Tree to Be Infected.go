package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func amountOfTime(root *TreeNode, start int) int {
	var st *TreeNode
	var parents = map[*TreeNode]*TreeNode{} // 父节点
	var dfs func(*TreeNode, *TreeNode)
	dfs = func(node, pa *TreeNode) {
		if node == nil {
			return
		}
		if node.Val == start {
			st = node
		}
		parents[node] = pa
		dfs(node.Left, node)
		dfs(node.Right, node)
	}
	dfs(root, nil)
	var ans, vis = -1, map[*TreeNode]bool{nil: true, st: true}
	for q := []*TreeNode{st}; len(q) > 0; ans++ {
		var tmp = q
		q = nil
		for _, node := range tmp {
			if node != nil {
				if !vis[node.Left] {
					vis[node.Left] = true
					q = append(q, node.Left)
				}
				if !vis[node.Right] {
					vis[node.Right] = true
					q = append(q, node.Right)
				}
				if p := parents[node]; !vis[p] {
					vis[p] = true
					q = append(q, p)
				}
			}
		}
	}
	return ans
}
