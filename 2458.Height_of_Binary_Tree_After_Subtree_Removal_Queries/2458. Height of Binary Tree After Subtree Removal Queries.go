package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeQueries(root *TreeNode, queries []int) []int {
	height := make(map[*TreeNode]int)
	var getHeight func(node *TreeNode) int
	getHeight = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		height[node] = max(getHeight(node.Left), getHeight(node.Right)) + 1
		return height[node]
	}
	getHeight(root)
	heightAfterRemoval := make([]int, len(height)+1)
	var dfs func(node *TreeNode, depth, restH int)
	dfs = func(node *TreeNode, depth, restH int) {
		if node == nil {
			return
		}
		depth++
		heightAfterRemoval[node.Val] = restH
		dfs(node.Left, depth, max(restH, depth+height[node.Right]))
		dfs(node.Right, depth, max(restH, depth+height[node.Left]))
	}
	dfs(root, -1, 0)
	for i, query := range queries {
		queries[i] = heightAfterRemoval[query]
	}
	return queries
}
