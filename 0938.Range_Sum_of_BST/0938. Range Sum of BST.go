package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归写法1：
//func rangeSumBST(root *TreeNode, low, high int) int {
//	if root == nil {
//		return 0
//	}
//	if root.Val > high {
//		return rangeSumBST(root.Left, low, high)
//	}
//	if root.Val < low {
//		return rangeSumBST(root.Right, low, high)
//	}
//	return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
//}

// 递归写法2：
func rangeSumBST(root *TreeNode, low int, high int) (ans int) {
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Val < low {
			dfs(node.Right)
		} else if node.Val > high {
			dfs(node.Left)
		} else {
			ans += node.Val
			dfs(node.Left)
			dfs(node.Right)
		}
	}
	dfs(root)
	return
}
