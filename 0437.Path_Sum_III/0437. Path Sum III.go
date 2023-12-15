package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) (ans int) {
	preSum := map[int]int{0: 1}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, curr int) {
		if node == nil {
			return
		}
		curr += node.Val              // 从根节点到当前节点的和
		ans += preSum[curr-targetSum] // 检查是否存在前缀和等于cur-targetSum
		preSum[curr]++                // 更新前缀和
		dfs(node.Left, curr)
		dfs(node.Right, curr)
		preSum[curr]-- // 退出时删除本节点
		return
	}
	dfs(root, 0)
	return
}
