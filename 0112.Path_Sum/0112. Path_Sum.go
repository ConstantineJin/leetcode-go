package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil { // 根节点为空，不存在根节点到叶子节点的路径
		return false
	} else if root.Left == nil && root.Right == nil { // 当前节点为叶子节点
		return root.Val == targetSum // 返回当前节点的值是否等于目标值
	} else {
		return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val) // 递归当且节点的左右孩子
	}
}
