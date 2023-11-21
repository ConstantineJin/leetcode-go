package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 数组
//func kthSmallest(root *TreeNode, k int) int {
//	var dfs func(node *TreeNode) []int // 将二叉搜索树转化为有序数组
//	dfs = func(node *TreeNode) []int {
//		if node == nil {
//			return nil
//		}
//		left, right := dfs(node.Left), dfs(node.Right)
//		return append(append(left, node.Val), right...) // 中序遍历
//	}
//	return dfs(root)[k-1]
//}

func kthSmallest(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		stack, root = stack[:len(stack)-1], stack[len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}
