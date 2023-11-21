package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	ans := math.MinInt
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left, right := dfs(node.Left), dfs(node.Right) // 递归左右子树，求其最大路径和
		innerMaxSum := left + node.Val + right         // 如果选择当前节点，那么最大路径和就是左右子树的最大路径和之和加上当前节点的值
		ans = max(ans, innerMaxSum)                    // 酌情更新答案
		return max(node.Val+max(left, right), 0)       // 要向上传递，那么当前节点的左右子树最多只能保留一颗，或者如果路径和为负就返回0（不选）
	}
	dfs(root)
	return ans
}
