package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minCameraCover(root *TreeNode) int {
	var dfs func(node *TreeNode) (a, b, c int) // a 表示 node 必须放置摄像头的情况下覆盖以 node 为根节点的子树所需要的摄像头数量；b 表示覆盖整棵子树需要的摄像头数目，无论 node 是否放置摄像头；c 表示覆盖两棵子树需要的摄像头数目，无论 node 本身是否被监控到
	dfs = func(node *TreeNode) (a, b, c int) {
		if node == nil {
			return math.MaxInt32, 0, 0
		}
		leftA, leftB, leftC := dfs(node.Left)
		rightA, rightB, rightC := dfs(node.Right)
		a = leftC + rightC + 1
		b = min(a, leftA+rightB, rightA+leftB)
		c = min(a, leftB+rightB)
		return
	}
	_, ans, _ := dfs(root)
	return ans
}
