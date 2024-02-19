package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) (ans int) {
	var dfs func(node *TreeNode) (sum, cnt int)
	dfs = func(node *TreeNode) (sum, cnt int) {
		if node == nil {
			return
		}
		var leftSum, leftCnt = dfs(node.Left)
		var rightSum, rightCnt = dfs(node.Right)
		sum, cnt = leftSum+rightSum+node.Val, leftCnt+rightCnt+1
		if node.Val == sum/cnt {
			ans++
		}
		return
	}
	dfs(root)
	return ans
}
